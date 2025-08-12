package errors

import "fmt"

var (
	ErrEventNotFound = fmt.Errorf("event not found")
)
