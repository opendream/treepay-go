package treepay

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strings"
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

type testBody struct {
	io.Reader

	body string
}

func (t testBody) Read(p []byte) (n int, err error) {
	return t.Reader.Read(p)
}

func (t testBody) Close() error {
	return nil
}

func NewTestBody(s string) testBody {
	return testBody{
		Reader: strings.NewReader(s),
	}
}

func TestAPIConnectionError_Error(t *testing.T) {
	e := APIConnectionError{treepayErr: &Error{
		HTTPResponse: &http.Response{
			StatusCode: http.StatusBadRequest,
			Body:       NewTestBody("something missing"),
		},
	}}
	assert.Equal(t, "Http400: something missing", e.Error())
}
