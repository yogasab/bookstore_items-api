package rest_errors_utils

import (
	"encoding/json"
	"errors"
	"net/http"
)

type RestErrors struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewErrFromBytes(bytes []byte) (*RestErrors, error) {
	var restErrors RestErrors
	err := json.Unmarshal(bytes, &restErrors)
	if err != nil {
		return nil, errors.New("invalid json data")
	}
	return &restErrors, nil
}

func NewGeneralError(code int, message string, status string, data interface{}) *RestErrors {
	generalError := RestErrors{
		Code:    code,
		Status:  status,
		Message: message,
		Data:    data,
	}
	return &generalError
}
func NewBadRequestError(message string) *RestErrors {
	badRequestError := RestErrors{
		Code:    http.StatusBadRequest,
		Status:  "failed",
		Message: message,
		Data:    nil,
	}
	return &badRequestError
}

func NewNotFoundError(message string) *RestErrors {
	notFoundError := RestErrors{
		Code:    http.StatusNotFound,
		Status:  "failed",
		Message: message,
		Data:    nil,
	}
	return &notFoundError
}

func NewUnauthorizedError(message string) *RestErrors {
	unauthorizedError := RestErrors{
		Code:    http.StatusUnauthorized,
		Status:  "failed",
		Message: message,
		Data:    nil,
	}
	return &unauthorizedError
}

func NewInternalServerError(message string) *RestErrors {
	internalServerError := RestErrors{
		Code:    http.StatusInternalServerError,
		Status:  "error",
		Message: message,
		Data:    nil,
	}
	return &internalServerError
}
