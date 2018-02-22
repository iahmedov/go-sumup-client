package transactions

//go:generate stringer -type EStatus,EPaymentType,ETransactionType -output enums.go

import (
	"time"
	"github.com/iahmedov/go-sumup-client/schemas/base"
)

type EStatus int
type EPaymentType int
type ETransactionType int

const (
	SUCCESSFUL EStatus = iota
	CANCELLED
	FAILED
)

const (
	CASH EPaymentType = iota
	POS
	ECOM
	BITCOIN
	BALANCE
)

const (
	PAYMENT ETransactionType = iota
	REFUND
	CHARGE_BACK
)

type FilterTransactionHistory struct {
	// SumUp transaction code.
	TransactionCode string

	// Ordering of results by timestamp. Possible values: ‘ascending’ (default), ‘descending’
	Order string

	// Maximum number of results per page
	Limit int64

	// Filter by user email
	UserEmails []string

	// Filter by geo coordinates: use CSVs of southwest_lng, southwest_lat, northeast_lng, northeast_lat
	// TODO: Support geocordinates by using GeoCoordinate struct and CSV wrapper for its String() function
	GeoCoordinates []string

	// Filter by reader serial number
	Readers []string

	// Filter by status. Possible values: SUCCESSFUL, CANCELLED, FAILED
	Statuses []EStatus

    // Filter by payment type. Possible values: CASH, POS, ECOM, BITCOIN, BALANCE
	PaymentTypes []EPaymentType

    // Filter by event type. Possible values: PAYMENT, REFUND, CHARGE_BACK
	Types []ETransactionType

    // Return only transactions that have been updated after the given timestamp. Format: ISO8601 (xmlschema)
	ChangesSince time.Time

	// Return only transactions that occurred before the given timestamp. Format: ISO8601 (xmlschema)
	NewestTime time.Time

    // Return only transactions that occurred after the given event ID
	NewestRef string

    // Whether to include “newest_ref” in result (default: false)
	NewestIncl bool

    // Return only transactions that occurred after the given timestamp. Format: ISO8601 (xmlschema)
	OldestTime time.Time

    // Return only transactions that occurred before the given event ID
	OldestRef string

	// Whether to include “oldest_ref” in result (default: false)
	OldestIncl bool
}

func (f *FilterTransactionHistory) KV() map[string]interface{} {
	return map[string]interface{}{}
}


var _ base.Filter = (*FilterTransactionHistory)(nil)