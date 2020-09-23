package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var (

	ErrInternal = &Error{
		Code : http.StatusInternalServerError,
		Message: "went wrong",
	}

	// ErrUnprocessableEntity HTTP 422
	ErrUnprocessableEntity = &Error{
		Code:    http.StatusUnprocessableEntity,
		Message: "Unprocessable Entity",
	}
	// ErrBadRequest HTTP 400
	ErrBadRequest = &Error{
		Code:    http.StatusBadRequest,
		Message: "Error invalid argument",
	}
	// ErrEventNotFound HTTP 404
	ErrEventNotFound = &Error{
		Code:    http.StatusNotFound,
		Message: "Event not found",
	}
	// ErrObjectIsRequired HTTP 400
	ErrObjectIsRequired = &Error{
		Code:    http.StatusBadRequest,
		Message: "Request object should be provided",
	}
	// ErrValidEventIDIsRequired HTTP 400
	ErrValidEventIDIsRequired = &Error{
		Code:    http.StatusBadRequest,
		Message: "A valid event id is required",
	}
	// ErrEventTimingIsRequired HTTP 400
	ErrEventTimingIsRequired = &Error{
		Code:    http.StatusBadRequest,
		Message: "Event start time and end time should be provided",
	}
	// ErrInvalidLimit HTTP 400
	ErrInvalidLimit = &Error{
		Code:    http.StatusBadRequest,
		Message: "Limit should be an integral value",
	}
	// ErrInvalidTimeFormat HTTP 400
	ErrInvalidTimeFormat = &Error{
		Code:    http.StatusBadRequest,
		Message: "Time Should be passed in RFC3339 Format: " + time.RFC3339,
	}

)


type Error struct {
	Code int
	Message string
}

func (err *Error) Error() string {
	if err == nil {
		return  ""
	}
	return fmt.Sprintf("error: code=%s message=%s", http.StatusText(err.Code), err.Message)
}

func (err *Error) String() string {
	return err.String()
}

func ( err *Error) JSON() [] byte {
	if err == nil {
		return  [] byte("{}")
	}
	res, _ := json.Marshal(err)
	return res
}

func  (err *Error)StatusCode() int {
	if err == nil {
		return http.StatusOK
	}
	return  err.Code


}

