package marvel_test

import (
	"fmt"
	"testing"

	"encoding/json"

	"github.com/dustinrc/marvel"
	"github.com/stretchr/testify/assert"
)

type mockAuth struct{}

func (ma *mockAuth) Auth() *marvel.AuthParams {
	return &marvel.AuthParams{
		Timestamp: "a",
		PublicKey: "b",
		Hash:      "c",
	}
}

func TestNewClient(t *testing.T) {
	auth := &mockAuth{}
	c := marvel.NewClient(auth)

	req, err := c.Request()

	assert.Nil(t, err)
	assert.Equal(t, "GET", req.Method)
	assert.Equal(t, "https://gateway.marvel.com/v1/public?apikey=b&hash=c&ts=a",
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
