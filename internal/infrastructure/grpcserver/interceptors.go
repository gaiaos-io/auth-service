package grpcserver

import (
	"context"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

func UnaryLoggingInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		start := time.Now()
		resp, err = handler(ctx, req)
		elapsed := time.Since(start)

		var peerAddr string
		if p, ok := peer.FromContext(ctx); ok {
			peerAddr = p.Addr.String()
		}

		code := codes.OK
		if err != nil {
			if st, ok := status.FromError(err); ok {
				code = st.Code()
			} else {
				code = codes.Unknown
			}
		}

		fields := []zap.Field{
			zap.String("grpc.method", info.FullMethod),
			zap.String("grpc.code", code.String()),
			zap.Duration("grpc.duration", elapsed),
		}
		if peerAddr != "" {
			fields = append(fields, zap.String("net.peer.ip", peerAddr))
		}

		if err != nil {
			if code >= codes.InvalidArgument && code < codes.Internal {
				// Expected client errors
				logger.Warn("gRPC request failed", append(fields, zap.Error(err))...)
			} else {
				// Server-side or unknown errors
				logger.Error("gRPC request failed", append(fields, zap.Error(err))...)
			}
		} else {
			logger.Info("gRPC request completed", fields...)
		}

		return resp, err
	}
}
