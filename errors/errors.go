package errors

import (
	"github.com/gorilla/rpc/v2/json2"
)

// HTTPForceJSONError type to pass json error to http handler
type HTTPForceJSONError struct {
	Error json2.Error
}

var (
	ErrInvalidArgument = &json2.Error{
		Code:    -32602,
		Message: "Invalid arguments",
	}
)
