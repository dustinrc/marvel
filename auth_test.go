package marvel_test

import (
	"testing"

	"github.com/dustinrc/marvel"
	"github.com/stretchr/testify/assert"
)

func TestServerSideAuth(t *testing.T) {
	auth := marvel.NewServerSideAuth("1234", "abcd")
	auth.Timestamper(func() string { return "1" })

	expected := &marvel.AuthParams{
		Timestamp: "1",
		PublicKey: "1234",
		Hash:      "ffd275c5130566a2916217b101f26150",
	}
	actual := auth.Auth()

	assert.Equal(t, expected, actual)
}

func TestClientSideAuth(t *testing.T) {
	auth := marvel.NewClientSideAuth("1234")

	expected := &marvel.AuthParams{
		PublicKey: "1234",
	}
	actual := auth.Auth()

	assert.Equal(t, expected, actual)
}
