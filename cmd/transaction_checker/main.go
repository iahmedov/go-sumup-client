package main

import (
	"time"
	sumup "github.com/iahmedov/go-sumup-client"
	st "github.com/iahmedov/go-sumup-client/schemas/transactions"
)

func main() {
	token := sumup.Token{
		Value: "1",
		ValidUntil: time.Now(),
	}
	refresh := token

	client, err := sumup.NewClient(token, refresh, nil)
	if err != nil {
		panic(err.Error())
	}

	enable := true
	client.AutoRefreshToken(enable)

	var txs *st.TransactionsHistory = client.Transactions.History(st.FilterTransactionHistory{})
	_ = txs
}