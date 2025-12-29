package grpcserver

import (
	"net"

	"google.golang.org/grpc"

	authpb "github.com/gaiaos-io/auth-service/proto/v1"
)

type Server struct {
	grpcServer *grpc.Server
	listener   net.Listener
}

func NewServer(addr string, authService authpb.AuthServiceServer, interceptors ...grpc.UnaryServerInterceptor) (*Server, error) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	opts := []grpc.ServerOption{}
	if len(interceptors) > 0 {
		opts = append(opts, grpc.ChainUnaryInterceptor(interceptors...))
	}

	grpcServer := grpc.NewServer(opts...)

	authpb.RegisterAuthServiceServer(grpcServer, authService)

	return &Server{
		grpcServer: grpcServer,
		listener:   lis,
	}, nil
}

func (server *Server) Start() error {
	return server.grpcServer.Serve(server.listener)
}

func (server *Server) Stop() {
	server.grpcServer.GracefulStop()
}
