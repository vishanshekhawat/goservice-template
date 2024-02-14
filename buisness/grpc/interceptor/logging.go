package interceptor

import (
	"context"
	"log"

	grpcF "github.com/vishn007/go-service-template/foundation/grpc"
	"google.golang.org/grpc"
)

// gRPC loggingInterceptor which helps to log
func LoggingInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ any, err error) {
		log.Printf("Received request: %v", req)
		v := grpcF.GetValues(ctx)
		resp, err := handler(ctx, req)
		log.Printf("Response: %v, %s,%v", resp, v.TraceID, err)
		return resp, err
	}
}
