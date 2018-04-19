package treepay

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestError_Error(t *testing.T) {
	e := Error{Code: "A101", Message: "Invalid request data"}
	assert.Equal(t, "A101: Invalid request data", e.Error())
}

func TestAPIError_Error(t *testing.T) {
	e := APIError{treepayErr: &Error{Code: "A101", Message: "Invalid request data"}}
	assert.Equal(t, "A101: Invalid request data", e.Error())
}
