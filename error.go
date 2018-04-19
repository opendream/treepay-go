package treepay

import (
	"fmt"
	"net/http"
)

type Error struct {
	Code    string `json:"res_cd"`
	Message string `json:"res_msg"`

	HTTPResponse *http.Response `json:"-"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

type APIError struct {
	treepayErr *Error
}

func (e *APIError) Error() string {
	return e.treepayErr.Error()
}
