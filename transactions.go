package client

import (
	"net/http"
	st "github.com/iahmedov/go-sumup-client/schemas/transactions"
)

type TransactionsService interface {
	History(filter st.FilterTransactionHistory) *st.TransactionsHistory
}

type sumupTransactionsService struct {
	// http client or kind of wrapper api? or transport? or some kind of client?
	httpClient *http.Client
}

func newTransactionsService(httpClient *http.Client) TransactionsService {
	return &sumupTransactionsService{
		httpClient: httpClient,
	}
}

func (s *sumupTransactionsService) History(filter st.FilterTransactionHistory) *st.TransactionsHistory {
	return nil
}

var _ TransactionsService = (*sumupTransactionsService)(nil)