package client

import (
	"context"
	"net/http"
	"time"

	sb "github.com/iahmedov/go-sumup-client/schemas/base"
)

const (
	THRESHOLD_MIN_TOKEN_VALIDITY = time.Second * 10
)

var (
	ENDPOINT = "https://api.sumup.com"
)

type Client struct {
	Transactions TransactionsService

	httpClient           *http.Client
	endpoint             string
	ctxAutoRefresh       context.Context
	ctxAutoRefreshCancel context.CancelFunc
}

func NewClient(access, refresh string, httpClient *http.Client) (*Client, error) {
	accessToken := sb.Token{
		Value:      access,
		ValidUntil: time.Now().Add(-THRESHOLD_MIN_TOKEN_VALIDITY),
	}
	refreshToken := sb.Token{
		Value:      refresh,
		ValidUntil: time.Now().Add(-THRESHOLD_MIN_TOKEN_VALIDITY),
	}

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	var baseTransport http.RoundTripper
	if httpClient.Transport == nil {
		baseTransport = http.DefaultTransport
	}

	t := &Transport{
		Base: baseTransport,
		source: &LockableTokenSource{
			access:  accessToken,
			refresh: refreshToken,
		},
		auth: BearerTokenAuth,
	}
	httpClient.Transport = t

	c := &Client{
		// public facing services/modules
		Transactions: newTransactionsService(httpClient, ENDPOINT+"/v0.1"),

		ctxAutoRefreshCancel: nil,
		httpClient:           httpClient,
	}

	// err := c.refreshToken()
	// if err != nil {
	// 	return nil, err
	// }

	return c, nil
}

// AutoRefreshToken enables auto refresh mode for tokens
func (c *Client) AutoRefreshToken(enable bool) context.CancelFunc {
	// TODO: what should we do when token expired or invalidated from server? should we respond to this case and try to
	// get another token?
	enabled := c.ctxAutoRefreshCancel != nil
	if enabled == enable {
		return c.ctxAutoRefreshCancel
	}

	if enable {
		c.ctxAutoRefresh, c.ctxAutoRefreshCancel = context.WithCancel(context.Background())

		go func(ctx context.Context) {
			ts := c.tokenSource()
			for {
				token, _ := ts.AccessToken()
				select {
				case <-ctx.Done():
					break
				case <-time.After(time.Until(token.ValidUntil) - THRESHOLD_MIN_TOKEN_VALIDITY):
					if time.Until(token.ValidUntil) > THRESHOLD_MIN_TOKEN_VALIDITY {
						// probably updated manually.
						// when token invalidated on the server side and client refreshed token by calling HandleError
						continue
					}

					c.refreshToken()
					break
				}
			}
		}(c.ctxAutoRefresh)

	} else {
		c.ctxAutoRefreshCancel()
		c.ctxAutoRefresh = nil
		c.ctxAutoRefreshCancel = nil
	}
	return c.ctxAutoRefreshCancel
}

func (c *Client) refreshToken() error {
	tokenSource := c.tokenSource()
	access, _ := tokenSource.AccessToken()
	refresh, _ := tokenSource.RefreshToken()

	noAuthClient := *c.httpClient
	transport, ok := noAuthClient.Transport.(*Transport)
	if !ok {
		panic("HTTP Client Transport must have *client.Transport type")
	}

	noAuthClient.Transport = &Transport{
		Base:   transport.Base,
		source: transport.source,
		auth:   NoAuth,
	}

	noAuthClient.Get("refresh token here")

	if ts, ok := tokenSource.(tokenSourceSetter); ok {
		ts.setAccessToken(&access)
		ts.setRefreshToken(&refresh)
		return nil
	}

	return ErrorImmutableTokenSource
}

func (c *Client) tokenSource() TokenSource {
	if t, ok := c.httpClient.Transport.(*Transport); ok {
		return t.source
	}
	return nil
}

func (c *Client) HandleError(err *sb.Error) error {
	if err.IsTokenExpired() {
		return c.refreshToken()
	}
	return nil
}
