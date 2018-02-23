package client

import (
	"fmt"
	"net/http"
)

type authConfig struct {
	token Token
	username string
	password string
	// do we need others?
}

type AuthSetter func(config authConfig, req *http.Request)

type Transport struct {
	Base http.RoundTripper
	source TokenSource
	auth AuthSetter
}

func (t *Transport) RoundTrip(r *http.Request) (*http.Response, error) {
	token, err := t.source.AccessToken()
	if err != nil {
		return nil, err
	}

	t.auth(authConfig{token:*token}, r)
	return t.base().RoundTrip(r)
}

func (t *Transport) base() http.RoundTripper {
	if t.Base == nil {
		t.Base = http.DefaultTransport
	}

	return t.Base
}

func BearerTokenAuth(config authConfig, req *http.Request) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.token))
}