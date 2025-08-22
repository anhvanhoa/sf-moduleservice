package grpcservice

import (
	"context"
	"time"

	"module-service/infrastructure/service/logger"

	"google.golang.org/grpc"
)

func LoggingInterceptor(log logger.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		start := time.Now()
		log.LogGRPCRequest(ctx, info.FullMethod, req)
		resp, err := handler(ctx, req)
		duration := time.Since(start)
		log.LogGRPC(ctx, info.FullMethod, req, resp, err, duration)
		log.LogGRPCResponse(ctx, info.FullMethod, resp, err, duration)
		return resp, err
	}
}
