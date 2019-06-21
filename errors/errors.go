package errors

import (
	"github.com/gorilla/rpc/v2/json2"
)

var (
	ErrInvalidArgument = &json2.Error{
		Code:    -32602,
		Message: "Invalid arguments",
	}

	ErrForbiddenURL = &json2.Error{
		Code:    -32603,
		Message: "Forbidden URL",
	}

	ErrInternalError = &json2.Error{
		Code:    -32604,
		Message: "Internal Server Error",
	}
)
