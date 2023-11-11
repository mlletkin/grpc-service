package middleware

import (
	"context"
	"time"

	"gitlab.ozon.dev/kavkazov/homework-8/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func KafkaLogging(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	l := logger.FromContext(ctx)
	ctx = logger.ToContext(ctx, l.With(zap.String("method", info.FullMethod)))
	logger.Infof(ctx, "%v", time.Now())
	return handler(ctx, req)
}
