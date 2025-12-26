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

func (*GrpcServer) PrimaryEmailRegister(ctx context.Context, req *authpb.PrimaryEmailRegisterRequest) (*authpb.AuthResponse, error) {
	panic("AuthService.PrimaryEmailRegister not implemented")
}

func (*GrpcServer) PrimaryEmailLogin(ctx context.Context, req *authpb.PrimaryEmailLoginRequest) (*authpb.AuthResponse, error) {
	panic("AuthService.PrimaryEmailLogin not implemented")
}

func (*GrpcServer) UpdatePrimaryEmail(ctx context.Context, req *authpb.UpdatePrimaryEmailRequest) (*emptypb.Empty, error) {
	panic("AuthService.UpdatePrimaryEmail not implemented")
}

func (*GrpcServer) VerifyPrimaryEmail(ctx context.Context, req *authpb.VerifyPrimaryEmailRequest) (*emptypb.Empty, error) {
	panic("AuthService.VerifyPrimaryEmail not implemented")
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

// User Info

func (*GrpcServer) GetMe(ctx context.Context, req *authpb.GetMeRequest) (*authpb.UserResponse, error) {
	panic("AuthService.GetMe not implemented")
}
