package main

import (
	sumup "github.com/iahmedov/go-sumup-client"
)

func main() {
	token := "1"
	refresh := "2"
	client := sumup.NewClient(token, refresh)
	enable := true
	client.TokenAutoRefresh(enable)
	client.Transactions()
}