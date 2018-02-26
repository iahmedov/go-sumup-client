package client

import (
	"fmt"
	"net/http"

	sb "github.com/iahmedov/go-sumup-client/schemas/base"
)

type authConfig struct {
	token    sb.Token
	username string
	password string
	// do we need others?
}

type AuthSetter func(config authConfig, req *http.Request)

type Transport struct {
	Base   http.RoundTripper
	source TokenSource
	auth   AuthSetter
}

func (t *Transport) RoundTrip(r *http.Request) (*http.Response, error) {
	token, err := t.source.AccessToken()
	if err != nil {
		return nil, err
	}

	t.auth(authConfig{token: token}, r)
	fmt.Println(r.Header)
	return t.base().RoundTrip(r)
}

func (t *Transport) base() http.RoundTripper {
	if t.Base == nil {
		t.Base = http.DefaultTransport
	}

	return t.Base
}

func BearerTokenAuth(config authConfig, req *http.Request) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.token.Value))
}

func NoAuth(config authConfig, req *http.Request) {
}
