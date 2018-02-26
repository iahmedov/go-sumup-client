package client

import (
	"net/http"

	sb "github.com/iahmedov/go-sumup-client/schemas/base"
	st "github.com/iahmedov/go-sumup-client/schemas/transactions"
)

type TransactionsService interface {
	History(filter st.FilterTransactionHistory) (*st.TransactionsHistory, *sb.Error, error)
}

type sumupTransactionsService struct {
	// http client or kind of wrapper api? or transport? or some kind of client?
	httpClient   *http.Client
	endpointPath string
}

func newTransactionsService(httpClient *http.Client, endpoint string) TransactionsService {
	return &sumupTransactionsService{
		httpClient:   httpClient,
		endpointPath: endpoint,
	}
}

func (s *sumupTransactionsService) History(filter st.FilterTransactionHistory) (*st.TransactionsHistory, *sb.Error, error) {
	resp, err := s.httpClient.Get(s.endpoint() + "/me/transactions/history")
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	var history st.TransactionsHistory
	var e sb.Error

	err = sb.DecodeInto(resp.Body, &history, &e)
	return &history, &e, err
}

func (s *sumupTransactionsService) endpoint() string {
	return s.endpointPath
}

var _ TransactionsService = (*sumupTransactionsService)(nil)
