package errors

import (
	"encoding/json"
	"fmt"
)

type HttpError struct {
	Cause error `json:"-"`
	Detail string `json:"detail"`
	Status int		`json:"-"`
}

func ( e *HttpError) Error() string {
	if e.Cause == nil {
		return e.Detail
	}
	return e.Detail + ":" + e.Cause.Error()
}

func (e *HttpError) ResponseBody() ([] byte , error) {

	body , err := json.Marshal(e)
	if err != nil {
		return  nil , fmt.Errorf("Error while parsing response body: %v", err)
	}


	return body , nil
}

func NewHttpError(err error, status int, detail string) *HttpError {
	return &HttpError{
		Cause:  err,
		Detail: detail,
		Status: status,
	}
}
