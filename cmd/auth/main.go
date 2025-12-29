package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"go.uber.org/zap"

	"github.com/gaiaos-io/auth-service/internal/infrastructure/grpcserver"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}
	defer logger.Sync()

	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	if env != "production" {
		if err := godotenv.Load(); err != nil {
			logger.Warn("No .env file found")
		}
	}

	addr := os.Getenv("AUTH_GRPC_ADDR")
	if addr == "" {
		addr = ":50051"
	}

	authService := grpcserver.NewAuthServiceServer()

	loggingIntercepter := grpcserver.UnaryLoggingInterceptor(logger)

	server, err := grpcserver.NewServer(addr, authService, loggingIntercepter)
	if err != nil {
		logger.Fatal("failed to create gRPC server", zap.Error(err))
	}

	go func() {
		logger.Info("gRPC server starting", zap.String("addr", addr))
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
