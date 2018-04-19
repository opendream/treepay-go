package treepay

import (
	"fmt"
	"io/ioutil"
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

type APIConnectionError struct {
	treepayErr *Error
}

func (e *APIConnectionError) Error() string {
	defer e.treepayErr.HTTPResponse.Body.Close()
	body, _ := ioutil.ReadAll(e.treepayErr.HTTPResponse.Body)

	return fmt.Sprintf("Http%d: %s", e.treepayErr.HTTPResponse.StatusCode, body)
}
