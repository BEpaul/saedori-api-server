package errors

import (
	"log"
)

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
		log.Fatalf("%s : %v", message, args)
	} else {
		log.Fatalf("unknown error code: %d", code)
	}
	return nil
}
