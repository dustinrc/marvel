package marvel_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
	"github.com/dustinrc/marvel"
	"github.com/stretchr/testify/assert"
)

// testClient acts like a normal marvel.Client, but also has a go-vcr recorder
// associated with it.
type testClient struct {
	*marvel.Client
	rec *recorder.Recorder
}

// newTestClient creates a new marvel.Client and associates the cassette path
// with it. The API keys are set using the environment variables "MARVEL_PUBLIC_KEY"
// and "MARVEL_PRIVATE_KEY".
func newTestClient(t *testing.T, cassette string) *testClient {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("no caller information when determining file location")
	}
	rec, err := recorder.New(filepath.Join(path.Dir(filename), "fixtures", cassette))
	if err != nil {
		t.Fatal("could not open cassette", cassette)
	}

	recHttpClient := &http.Client{
		Transport: rec,
	}

	pubKey, pubOK := os.LookupEnv("MARVEL_PUBLIC_KEY")
	privKey, privOK := os.LookupEnv("MARVEL_PRIVATE_KEY")
	if !pubOK || !privOK {
		t.Fatal("environment variables MARVEL_PUBLIC_KEY and MARVEL_PRIVATE_KEY not set")
	}

	auth := marvel.NewServerSideAuth(pubKey, privKey)
	auth.Timestamper(func() string { return "1" })
	c := marvel.NewClient(auth, recHttpClient)

	return &testClient{c, rec}
}

// stopRecorder stops and closes the recorder associated with the testClient.
func (tc *testClient) stopRecorder() { tc.rec.Stop() }

// mockAuth exists solely for its Auth() implementation.
type mockAuth struct{}

// Auth implements the Authenticator interface. It guarantees the auth query
// parameters of '?apikey=b&hash=c&ts=a'
func (ma *mockAuth) Auth() *marvel.AuthParams {
	return &marvel.AuthParams{
		Timestamp: "a",
		PublicKey: "b",
		Hash:      "c",
	}
}

func TestNewClient(t *testing.T) {
	auth := &mockAuth{}
	c := marvel.NewClient(auth, nil)

	req, err := c.Request("", nil)

	assert.Nil(t, err)
	assert.Equal(t, "GET", req.Method)
	assert.Equal(t, "https://gateway.marvel.com/v1/public/?apikey=b&hash=c&ts=a",
		req.URL.String())
}

func TestAPIErrorUnmarshal(t *testing.T) {
	testCases := []struct {
		desc, jIn, eMsg string
	}{
		{
			desc: "APIError code is an integer",
			jIn:  `{"code": 404, "message": "something's not found"}`,
			eMsg: "marvel: 404 something's not found",
		},
		{
			desc: "APIError code is a string",
			jIn:  `{"code": "NotFound", "message": "still not found"}`,
			eMsg: "marvel: NotFound still not found",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ae := new(marvel.APIError)
			err := json.Unmarshal([]byte(tC.jIn), ae)
			assert.Nil(t, err)
			assert.Equal(t, tC.eMsg, fmt.Sprint(ae))
		})
	}
}
