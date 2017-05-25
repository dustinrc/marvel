package marvel_test

import (
	"testing"

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
