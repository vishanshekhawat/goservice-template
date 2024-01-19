package middleware

import (
	"context"
	"net/http"

	"github.com/vishn007/go-service-template/buisness/customerrors"
	"github.com/vishn007/go-service-template/foundation/logger"
	"github.com/vishn007/go-service-template/foundation/web"
)

// ErrorResponse is the form used for API responses from failures in the API.
type ErrorResponse struct {
	Error  string            `json:"error"`
	Fields map[string]string `json:"fields,omitempty"`
}

// Errors handles errors coming out of the call chain. It detects normal
// application errors which are used to respond to the client in a uniform way.
// Unexpected errors (status >= 500) are logged.
func Errors(log *logger.Logger) web.Middleware {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			err := handler(ctx, w, r)
			if err != nil {

				var er ErrorResponse
				var status int

				switch {

				case customerrors.IsRequestError(err):
					reqErr := customerrors.GetRequestError(err)
					er = ErrorResponse{
						Error: reqErr.Error(),
					}
					status = reqErr.Status

				case customerrors.IsRateLimitError(err):
					reqErr := customerrors.GetRateLimitError(err)
					er = ErrorResponse{
						Error: reqErr.Error(),
					}
					status = reqErr.Status

				default:
					er = ErrorResponse{
						Error: http.StatusText(http.StatusInternalServerError),
					}
					status = http.StatusInternalServerError
				}

				if err := web.Respond(ctx, w, er, status); err != nil {
					return err
				}
			}

			return err
		}

		return h
	}

	return m
}
