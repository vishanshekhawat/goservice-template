package interceptor

import (
	"context"
	"log"

	grpcF "github.com/vishn007/go-service-template/foundation/grpc"
	"google.golang.org/grpc"
)

// gRPC loggingInterceptor which helps to log
func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	log.Printf("Received request: %v", req)
	v := grpcF.GetValues(ctx)
	resp, err := handler(ctx, req)
	log.Printf("Response: %v, %s", resp, v.TraceID)
	return resp, err
}
