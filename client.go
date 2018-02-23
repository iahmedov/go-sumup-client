package client

type Client struct {
	Transactions *TransactionsService

	transport *Transport
}

func NewClient(access, refresh Token) *Client {
	transport := &Transport{
		Base:   nil,
		source: &LockableTokenSource{},
		auth: newAuther(config),
	}
	c.transport = transport
	// &http.Client{Transport: transport}

	c := &Client{}
	return c
}
