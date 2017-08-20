package marvel

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"
)

// TimestampFunc provides a value to use as Timestamp in AuthParams.
type TimestampFunc func() string

func defaultTimestamper() string {
	return strconv.FormatInt(time.Now().UTC().UnixNano(), 16)
}

// AuthParams are the query parameter values provided to the API for authentication.
// PublicKey must be provided for client side authentication. Server side
// authentication requires Timestamp, PublicKey and Hash.
type AuthParams struct {
	Timestamp string `url:"ts,omitempty"`
	PublicKey string `url:"apikey"`
	Hash      string `url:"hash,omitempty"`
}

// Authenticator is the interface for providing AuthParams, namely to a Client.
type Authenticator interface {
	Auth() *AuthParams
}

// ServerSideAuth holds the API keys and timestamp function necessary for server
// side authentication.
type ServerSideAuth struct {
	pubKey  string
	privKey string
	tsFunc  TimestampFunc
}

// NewServerSideAuth returns a ServerSideAuth which uses a default TimestampFunc
// based on the current epoch.
func NewServerSideAuth(publicKey, privateKey string) *ServerSideAuth {
	return &ServerSideAuth{
		pubKey:  publicKey,
		privKey: privateKey,
		tsFunc:  defaultTimestamper,
	}
}

// Auth implements the Authenticator interface.
func (ssa *ServerSideAuth) Auth() *AuthParams {
	ts := ssa.tsFunc()
	hasher := md5.New()
	hasher.Write([]byte(ts + ssa.privKey + ssa.pubKey))
	hash := hex.EncodeToString(hasher.Sum(nil))

	return &AuthParams{
		Timestamp: ts,
		PublicKey: ssa.pubKey,
		Hash:      hash,
	}
}

// Timestamper replaces the default TimestampFunc with the one provided as an
// argument.
func (ssa *ServerSideAuth) Timestamper(timestamper TimestampFunc) {
	ssa.tsFunc = timestamper
}

// ClientSideAuth holds the public API key necessary for client side authentication,
// enabling this package to be used with GopherJS. Client side requests (i.e.,
// browser-based) must originate from a pre-authorized web site or browser extension
// URL. Authorized sites and extensions are configurable from the Marvel developer portal.
type ClientSideAuth struct {
	pubKey string
}

// NewClientSideAuth returns a ClientSideAuth using the provided public API key.
func NewClientSideAuth(publicKey string) *ClientSideAuth {
	return &ClientSideAuth{
		pubKey: publicKey,
	}
}

// Auth implements the Authenticator interface.
func (csa *ClientSideAuth) Auth() *AuthParams {
	return &AuthParams{
		PublicKey: csa.pubKey,
	}
}
