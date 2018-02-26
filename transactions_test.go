package client

import (
	"strings"
	"testing"

	sb "github.com/iahmedov/go-sumup-client/schemas/base"
	st "github.com/iahmedov/go-sumup-client/schemas/transactions"
)

func TestTransactionHistoryParse(t *testing.T) {
	var history st.TransactionsHistory
	var e sb.Error

	transactionsHistory := `{
		"items": [
		{
		  "id": "...",
		  "transaction_id": "...",
		  "user": "...",
		  "type": "...",
		  "status": "...",
		  "timestamp": "...",
		  "currency": "...",
		  "amount": 1,
		  "transaction_code": "...",
		  "product_summary": "...",
		  "installments_count": 1,
		  "payment_type": "...",
		  "card_type": "...",
		  "payouts_total": 1,
		  "payouts_received": 1,
		  "payout_date": "...",
		  "payout_plan": "...",
		  "payout_type": "BANK_ACCOUNT"
		}
	  ],
	  "links": [
		{
		  "rel": "...",
		  "href": "...",
		  "type": "..."
		}
	  ]
	  }`

	err := sb.DecodeInto(strings.NewReader(transactionsHistory), &history, &e)
	if err != nil {
		t.Fatalf("Failed to parse transactions.history: %s", err.Error())
	}

	if len(history.Transactions) != 1 {
		t.Fatalf("failed to parse transaction history items")
	}

	if history.Transactions[0].PayoutType != "BANK_ACCOUNT" {
		t.Fatalf("failed to parse correctly")
	}
}
