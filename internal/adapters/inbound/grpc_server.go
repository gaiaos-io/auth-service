package inbound

import (
	"github.com/gaiaos-io/auth-service/proto/v1"

	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ authpb.AuthServiceServer = (*GrpcServer)(nil)

type GrpcServer struct {
	authpb.UnimplementedAuthServiceServer
}

// Email-Based Authentication

func (*GrpcServer) EmailRegister(ctx context.Context, req *authpb.EmailRegisterRequest) (*authpb.AuthResponse, error) {
	panic("AuthService.EmailRegister not implemented")
}

func (*GrpcServer) EmailLogin(ctx context.Context, req *authpb.EmailLoginRequest) (*authpb.AuthResponse, error) {
	panic("AuthService.EmailLogin not implemented")
}

// OAuth-Based Authentication

func (*GrpcServer) StartOAuth(ctx context.Context, req *authpb.StartOAuthRequest) (*authpb.StartOAuthResponse, error) {
	panic("AuthService.StartOAuth not implemented")
}

func (*GrpcServer) FinishOAuth(ctx context.Context, req *authpb.FinishOAuthRequest) (*authpb.AuthResponse, error) {
	panic("AuthService.FinishOAuth not implemented")
}

func (*GrpcServer) UnlinkOAuth(ctx context.Context, req *authpb.UnlinkOAuthRequest) (*emptypb.Empty, error) {
	panic("AuthService.UnlinkOAuth not implemented")
}

// Token & Session Management

func (*GrpcServer) RefreshToken(ctx context.Context, req *authpb.RefreshTokenRequest) (*authpb.AuthResponse, error) {
	panic("AuthService.RefreshToken not implemented")
}

func (*GrpcServer) Logout(ctx context.Context, req *authpb.LogoutRequest) (*emptypb.Empty, error) {
	panic("AuthService.Logout not implemented")
}

// User Management

func (*GrpcServer) UpdateContactEmail(ctx context.Context, req *authpb.UpdateContactEmailRequest) (*emptypb.Empty, error) {
	panic("AuthService.UpdateContactEmail not implemented")
}

func (*GrpcServer) VerifyContactEmail(ctx context.Context, req *authpb.VerifyContactEmailRequest) (*emptypb.Empty, error) {
	panic("AuthService.VerifyContactEmail not implemented")
}

func (*GrpcServer) GetMe(ctx context.Context, req *authpb.GetMeRequest) (*authpb.UserResponse, error) {
	panic("AuthService.GetMe not implemented")
}
