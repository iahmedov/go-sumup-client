package client

import "fmt"

var (
	ErrorImmutableTokenSource = fmt.Errorf("Update issued to immutable TokenStore")
)