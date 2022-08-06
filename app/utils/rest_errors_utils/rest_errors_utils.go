package rest_errors_utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type RestErrors interface {
	Code() int
	Status() string
	Message() string
	Data() []interface{}
}

type restErrors struct {
	code    int           `json:"code"`
	status  string        `json:"status"`
	message string        `json:"message"`
	data    []interface{} `json:"data"`
}

func NewRestErrors(code int, status string, message string, data []interface{}) RestErrors {
	return restErrors{
		code:    code,
		status:  status,
		message: message,
		data:    data,
	}
}

func (re restErrors) Code() int {
	return re.code
}

func (re restErrors) Status() string {
	return re.status
}

func (re restErrors) Message() string {
	return fmt.Sprintf("message: %s - code: %d - status: %s - data: %v", re.message, re.code, re.status, re.data)
}

func (re restErrors) Data() []interface{} {
	return re.data
}

func NewErrFromBytes(bytes []byte) (*RestErrors, error) {
	var restErrors RestErrors
	err := json.Unmarshal(bytes, &restErrors)
	if err != nil {
		return nil, errors.New("invalid json data")
	}
	return &restErrors, nil
}

func NewGeneralError(code int, message string, status string, data []interface{}) RestErrors {
	generalError := restErrors{
		code:    code,
		status:  status,
		message: message,
		data:    data,
	}
	return generalError
}
func NewBadRequestError(message string) RestErrors {
	badRequestError := restErrors{
		code:    http.StatusBadRequest,
		status:  "failed",
		message: message,
		data:    nil,
	}
	return badRequestError
}

func NewNotFoundError(message string) RestErrors {
	notFoundError := restErrors{
		code:    http.StatusNotFound,
		status:  "failed",
		message: message,
		data:    nil,
	}
	return notFoundError
}

func NewUnauthorizedError(message string) RestErrors {
	unauthorizedError := restErrors{
		code:    http.StatusUnauthorized,
		status:  "failed",
		message: message,
		data:    nil,
	}
	return unauthorizedError
}

func NewInternalServerError(message string, err error) RestErrors {
	internalServerError := restErrors{
		code:    http.StatusInternalServerError,
		status:  "failed",
		message: message,
		data:    nil,
	}
	if err != nil {
		internalServerError.data = append(internalServerError.data, err.Error())
	}
	return internalServerError
}
