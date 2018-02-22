package base

// Filter used in filtering queries
type Filter interface {
	KV() map[string]interface{}
}

// URLEncodedQuery used as a URL query ('?') use this method
func URLEncodedQuery(f Filter) string {
	return ""
}

// FormDataEncoded used inside the POST query as form-data
func FormDataEncoded(f Filter) string {
	return ""
}