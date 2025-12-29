package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/gaiaos-io/auth-service/internal/config"
	"github.com/gaiaos-io/auth-service/internal/infrastructure/grpcserver"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}
	defer logger.Sync()

	cfg := config.LoadConfig(logger)

	logger.Info("starting service", zap.String("ENV", cfg.Env))

	authService := grpcserver.NewAuthServiceServer()

	loggingIntercepter := grpcserver.UnaryLoggingInterceptor(logger)

	server, err := grpcserver.NewServer(cfg.GRPCAddr, authService, loggingIntercepter)
	if err != nil {
		logger.Fatal("failed to create gRPC server", zap.Error(err))
	}

	go func() {
		logger.Info("gRPC server starting", zap.String("addr", cfg.GRPCAddr))
		if err := server.Start(); err != nil {
			logger.Fatal("gRPC server failed", zap.Error(err))
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown

	logger.Info("shutting down gRPC server...")
	server.Stop()
	logger.Info("server stopped gracefully")
}
