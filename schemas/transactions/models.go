package transactions

////////////////////////////////////////////////
//
//  generated using transaction_detail.json
//
////////////////////////////////////////////////

// Card Card information
type Card struct {
	Last4Digits string `json:"last_4_digits,omitempty"`
	Type        string `json:"type,omitempty"`
}

// ElvAccount Elv bank account
type ElvAccount struct {
	Iban        string `json:"iban,omitempty"`
	Last4Digits string `json:"last_4_digits,omitempty"`
	SequenceNo  string `json:"sequence_no,omitempty"`
	SortCode    string `json:"sort_code,omitempty"`
}

// Event A list with transaction events
type Event struct {
	Amount        float64 `json:"amount,omitempty"`
	FeeAmount     float64 `json:"fee_amount,omitempty"`
	Id            float64 `json:"id,omitempty"`
	ReceiptNo     string  `json:"receipt_no,omitempty"`
	Status        string  `json:"status,omitempty"`
	Timestamp     string  `json:"timestamp,omitempty"`
	TransactionId string  `json:"transaction_id,omitempty"`
	Type          string  `json:"type,omitempty"`
}

// Link A list with links to related resources
type Link struct {
	Href string `json:"href,omitempty"`
	Rel  string `json:"rel,omitempty"`
	Type string `json:"type,omitempty"`
}

// Location Location information
type Location struct {
	HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`
	Lat                float64 `json:"lat,omitempty"`
	Lon                float64 `json:"lon,omitempty"`
}

// Product Purchase products
type Product struct {
	Description     string  `json:"description,omitempty"`
	Name            string  `json:"name,omitempty"`
	Price           float64 `json:"price,omitempty"`
	PriceWithVat    float64 `json:"price_with_vat,omitempty"`
	Quantity        float64 `json:"quantity,omitempty"`
	SingleVatAmount float64 `json:"single_vat_amount,omitempty"`
	TotalPrice      float64 `json:"total_price,omitempty"`
	TotalWithVat    float64 `json:"total_with_vat,omitempty"`
	VatAmount       float64 `json:"vat_amount,omitempty"`
	VatRate         float64 `json:"vat_rate,omitempty"`
}

// Transaction <p>&nbsp;</p>
type Transaction struct {
	Amount               float64           `json:"amount,omitempty"`
	AuthCode             string            `json:"auth_code,omitempty"`
	Card                 *Card             `json:"card,omitempty"`
	Currency             string            `json:"currency,omitempty"`
	ElvAccount           *ElvAccount       `json:"elv_account,omitempty"`
	EntryMode            string            `json:"entry_mode,omitempty"`
	Events               *Event            `json:"events,omitempty"`
	ForeignTransactionId string            `json:"foreign_transaction_id,omitempty"`
	HorizontalAccuracy   float64           `json:"horizontal_accuracy,omitempty"`
	Id                   string            `json:"id,omitempty"`
	InstallmentsCount    float64           `json:"installments_count,omitempty"`
	InternalId           float64           `json:"internal_id,omitempty"`
	Lat                  float64           `json:"lat,omitempty"`
	Links                *Link             `json:"links,omitempty"`
	LocalTime            string            `json:"local_time,omitempty"`
	Location             *Location         `json:"location,omitempty"`
	Lon                  float64           `json:"lon,omitempty"`
	MerchantCode         string            `json:"merchant_code,omitempty"`
	PaymentType          string            `json:"payment_type,omitempty"`
	PayoutDate           string            `json:"payout_date,omitempty"`
	PayoutPlan           string            `json:"payout_plan,omitempty"`
	PayoutType           string            `json:"payout_type,omitempty"`
	PayoutsReceived      float64           `json:"payouts_received,omitempty"`
	PayoutsTotal         float64           `json:"payouts_total,omitempty"`
	ProcessAs            string            `json:"process_as,omitempty"`
	ProductSummary       string            `json:"product_summary,omitempty"`
	Products             *Product          `json:"products,omitempty"`
	SimplePaymentType    string            `json:"simple_payment_type,omitempty"`
	SimpleStatus         string            `json:"simple_status,omitempty"`
	Status               string            `json:"status,omitempty"`
	TaxEnabled           bool              `json:"tax_enabled,omitempty"`
	Timestamp            string            `json:"timestamp,omitempty"`
	TipAmount            float64           `json:"tip_amount,omitempty"`
	TransactionCode      string            `json:"transaction_code,omitempty"`
	TransactionEvents    *TransactionEvent `json:"transaction_events,omitempty"`
	Username             string            `json:"username,omitempty"`
	VatAmount            float64           `json:"vat_amount,omitempty"`
	VerificationMethod   string            `json:"verification_method,omitempty"`
}

// TransactionEvent A list of transaction events such as payouts, refunds and chargebacks
type TransactionEvent struct {
	Amount            float64 `json:"amount,omitempty"`
	Date              string  `json:"date,omitempty"`
	DueDate           string  `json:"due_date,omitempty"`
	EventType         string  `json:"event_type,omitempty"`
	Id                float64 `json:"id,omitempty"`
	InstallmentNumber float64 `json:"installment_number,omitempty"`
	Status            string  `json:"status,omitempty"`
	Timestamp         string  `json:"timestamp,omitempty"`
}

////////////////////////////////////////////////
//
//  generated using transaction_history.json
//
////////////////////////////////////////////////

// TransactionLite
type TransactionHistoryItem struct {
	Amount            float64 `json:"amount,omitempty"`
	CardType          string  `json:"card_type,omitempty"`
	Currency          string  `json:"currency,omitempty"`
	Id                string  `json:"id,omitempty"`
	InstallmentsCount float64 `json:"installments_count,omitempty"`
	PaymentType       string  `json:"payment_type,omitempty"`
	PayoutDate        string  `json:"payout_date,omitempty"`
	PayoutPlan        string  `json:"payout_plan,omitempty"`
	PayoutType        string  `json:"payout_type,omitempty"`
	PayoutsReceived   float64 `json:"payouts_received,omitempty"`
	PayoutsTotal      float64 `json:"payouts_total,omitempty"`
	ProductSummary    string  `json:"product_summary,omitempty"`
	Status            string  `json:"status,omitempty"`
	Timestamp         string  `json:"timestamp,omitempty"`
	TransactionCode   string  `json:"transaction_code,omitempty"`
	TransactionId     string  `json:"transaction_id,omitempty"`
	Type              string  `json:"type,omitempty"`
	User              string  `json:"user,omitempty"`
}

// TransactionsHistory
type TransactionsHistory struct {
	Links        []Link                   `json:"links,omitempty"`
	Transactions []TransactionHistoryItem `json:"items,omitempty"`
}
