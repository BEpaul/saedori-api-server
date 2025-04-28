package errors

import (
	"fmt"
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
		log.Println(message, args)
		return fmt.Errorf("%s : %v", message, args)
	} else {
		log.Println("unknown error code:", code)
		return fmt.Errorf("unknown error code: %d", code)
	}
}

func (e *Error) Error() string {
	if e.Code == 0 {
		log.Println(e.Message, e.Args)
		return fmt.Sprintf("%s : %v", e.Message, e.Args)
	}
	log.Println("unknown error code:", e.Code)
	return fmt.Sprintf("unknown error code: %d", e.Code)
}
