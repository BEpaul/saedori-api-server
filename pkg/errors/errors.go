package errors

import "fmt"

const (
	invalidParameter    = 1000
	internalServerError = 2000
	notFound            = 3000
	unauthorized        = 4000
	badRequest          = 5000
	conflict            = 6000
)

var errMessage = map[int64]string{
	invalidParameter:    "INVALID_PARAMETER",
	internalServerError: "INTERNAL_SERVER_ERROR",
	notFound:            "NOT_FOUND",
	unauthorized:        "UNAUTHORIZED",
	badRequest:          "BAD_REQUEST",
	conflict:            "CONFLICT",
}

func Errorf(code int64, args ...interface{}) error {
	if message, ok := errMessage[code]; ok {
		return fmt.Errorf("%s : %v", message, args)
	} else {
		return fmt.Errorf("unknown error code: %d", code)
	}
}
