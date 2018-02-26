package main

import (
	"fmt"

	sumup "github.com/iahmedov/go-sumup-client"
	st "github.com/iahmedov/go-sumup-client/schemas/transactions"
)

func main() {
	token := "1"
	refresh := "2"

	client, err := sumup.NewClient(token, refresh, nil)
	if err != nil {
		panic(err.Error())
	}

	enable := false
	cancelFunc := client.AutoRefreshToken(enable)
	defer func() {
		if cancelFunc != nil {
			cancelFunc()
		}
	}()

	txs, apiErr, err := client.Transactions.History(st.FilterTransactionHistory{})
	_ = txs

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(txs, apiErr)

	if apiErr != nil && apiErr.IsTokenExpired() {
		client.HandleError(apiErr)
	}
}
