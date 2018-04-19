package treepay

import (
	"context"
	"net/http"
)

type Params struct {
	Headers        http.Header
	PaymentRequest *PaymentRequest

	Context           context.Context
	ShouldSignRequest bool
	SecureKey         string
}
