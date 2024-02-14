package interceptor

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

// gRPC loggingInterceptor which helps to log
func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("Received request: %v", req)
	resp, err := handler(ctx, req)
	log.Printf("Response: %v", resp)
	return resp, err
}
