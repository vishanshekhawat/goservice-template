package interceptor

import (
	"context"
	"fmt"
	"runtime/debug"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// PanicInterceptor returns a new unary server interceptor for panic recovery.
func PanicInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ any, err error) {
		defer func() {
			if rec := recover(); rec != nil {
				trace := debug.Stack()
				p := fmt.Errorf("PANIC [%v] TRACE[%s]", rec, string(trace))
				err = status.Errorf(codes.Internal, "%s", p)
			}
		}()
		return handler(ctx, req)
	}
}
