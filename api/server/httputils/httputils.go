package httputils

import (
	"context"
	"github.com/vietnamz/prime-generator/errdefs"
	"net/http"
)

type APIVersionKey struct{}

type APIFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error


// MakeErrorHandler makes an HTTP handler that decodes a Error and
// returns it in the response.
func MakeErrorHandler( err error ) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		statusCode := errdefs.GetHTTPErrorStatusCode(err)
		http.Error(w, err.Error(), statusCode)
	}
}
