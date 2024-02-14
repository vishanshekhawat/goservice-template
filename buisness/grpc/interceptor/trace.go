package interceptor

import (
	"context"
	"time"

	"github.com/google/uuid"
	grpcF "github.com/vishn007/go-service-template/foundation/grpc"
	"google.golang.org/grpc"
)

// gRPC TraceInterceptor which helps to log
func TraceInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ any, err error) {

		requestID := uuid.NewString()

		v := grpcF.Values{
			TraceID:       uuid.NewString(),
			CoRealationID: requestID,
			Now:           time.Now().UTC(),
		}
		ctx = context.WithValue(ctx, grpcF.TraceKey, &v)
		resp, err := handler(ctx, req)
		return resp, err
	}
}
