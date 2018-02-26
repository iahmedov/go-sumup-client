package base

import "time"

type Token struct {
	Value      string
	ValidUntil time.Time
}

// Error struct
type Error struct {
	Code    string `json:"error_code"`
	Message string `json:"error_message"`
}

func (e *Error) IsTokenExpired() bool {
	return false
}

func (e *Error) HasError() bool {
	return len(e.Code) > 0
}
