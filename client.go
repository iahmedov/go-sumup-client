package client

import (
	"context"
	"net/http"
)

type Client struct {
	Transactions TransactionsService

	transport *Transport
	ctxAutoRefreshCancel context.CancelFunc
}

func NewClient(access, refresh Token, httpClient *http.Client) (*Client, error) {

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	var baseTransport http.RoundTripper
	if httpClient.Transport == nil {
		baseTransport = http.DefaultTransport
	}

	t := &Transport{
		Base:   baseTransport,
		source: &LockableTokenSource{
			access: access,
			refresh: refresh,
		},
		auth: BearerTokenAuth,
	}
	httpClient.Transport = t

	c := &Client{
		// public facing services/modules
		Transactions: newTransactionsService(httpClient),

		// internal
		ctxAutoRefreshCancel: nil,
		transport: t,
	}

	err := c.refreshToken()
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) AutoRefreshToken(enable bool) {
	enabled := c.ctxAutoRefreshCancel != nil
	if enabled == enable {
		return
	}

	if enable {
		ctx, cancelFunc := context.WithCancel(context.Background())
		c.ctxAutoRefreshCancel = cancelFunc

		go func(ctx context.Context) {
			for {
				select {
				case <-ctx.Done():
					break
				}
			}
		}(ctx)

	} else {
		c.ctxAutoRefreshCancel()
		c.ctxAutoRefreshCancel = nil
	}
}

func (c *Client) refreshToken() error {
	tokenSource := c.tokenSource()
	access, _ := tokenSource.AccessToken()
	refresh, _ := tokenSource.RefreshToken()

	if ts, ok := tokenSource.(tokenSourceSetter); ok {
		ts.setAccessToken(access)
		ts.setRefreshToken(refresh)
		return nil
	}

	return ErrorImmutableTokenSource 
}

func (c *Client) tokenSource() TokenSource {
	return c.transport.source
}