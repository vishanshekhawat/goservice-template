package interceptor

import (
	"context"
	"fmt"

	"github.com/vishn007/go-service-template/buisness/customerrors"
	"github.com/vishn007/go-service-template/buisness/validate"
	"github.com/vishn007/go-service-template/buisness/web/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ErrorInterceptor returns a new unary server interceptor for panic recovery.
func ErrorInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ any, err error) {

		resp, err := handler(ctx, req)
		if err != nil {
			fmt.Printf("%#v", err)

			switch {
			case validate.IsFieldErrors(err):

				err = status.Error(400, "Default error message for 400")
			case customerrors.IsRequestError(err):
				err = status.Error(400, "request error")
			case auth.IsAuthError(err):
				err = status.Error(400, "Auth error")
			case customerrors.IsRateLimitError(err):
				err = status.Error(400, "Rate Limit error")
			default:
				err = status.Error(codes.Internal, "internal server error")
			}
		}
		return resp, err
	}
}
