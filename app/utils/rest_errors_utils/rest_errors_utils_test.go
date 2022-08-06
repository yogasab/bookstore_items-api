package rest_errors_utils

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInternalServerError(t *testing.T) {
	restErrors := NewInternalServerError("internal server error message", errors.New("database error"))
	assert.NotNil(t, restErrors)
	assert.EqualValues(t, http.StatusInternalServerError, restErrors.Code())
	assert.EqualValues(t, "failed", restErrors.Status())
	assert.EqualValues(t, "message: internal server error message - code: 500 - status: failed - data: [database error]", restErrors.Message())

	assert.NotNil(t, restErrors.Data())
	assert.EqualValues(t, 1, len(restErrors.Data()))
	assert.EqualValues(t, "database error", restErrors.Data()[0])
}
func TestNewBadRequestError(t *testing.T) {
}
func TestNewNotFoundError(t *testing.T) {
}
func TestUnauthorizedError(t *testing.T) {
}
