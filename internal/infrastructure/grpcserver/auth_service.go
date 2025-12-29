package grpcserver

import (
	"github.com/gaiaos-io/auth-service/proto/v1"

	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ authpb.AuthServiceServer = (*AuthServiceServer)(nil)

type AuthServiceServer struct {
	authpb.UnimplementedAuthServiceServer
}

func NewAuthServiceServer() *AuthServiceServer {
	return &AuthServiceServer{}
}

// Email-Based Authentication

func (*Server) EmailRegister(ctx context.Context, req *authpb.EmailRegisterRequest) (*authpb.AuthResponse, error) {
	panic("AuthService.EmailRegister not implemented")
}

func (*Server) EmailLogin(ctx context.Context, req *authpb.EmailLoginRequest) (*authpb.AuthResponse, error) {
	panic("AuthService.EmailLogin not implemented")
}

// OAuth-Based Authentication

func (*Server) StartOAuth(ctx context.Context, req *authpb.StartOAuthRequest) (*authpb.StartOAuthResponse, error) {
	panic("AuthService.StartOAuth not implemented")
}

func (*Server) FinishOAuth(ctx context.Context, req *authpb.FinishOAuthRequest) (*authpb.AuthResponse, error) {
	panic("AuthService.FinishOAuth not implemented")
}

func (*Server) UnlinkOAuth(ctx context.Context, req *authpb.UnlinkOAuthRequest) (*emptypb.Empty, error) {
	panic("AuthService.UnlinkOAuth not implemented")
}

// Token & Session Management

func (*Server) RefreshToken(ctx context.Context, req *authpb.RefreshTokenRequest) (*authpb.AuthResponse, error) {
	panic("AuthService.RefreshToken not implemented")
}

func (*Server) Logout(ctx context.Context, req *authpb.LogoutRequest) (*emptypb.Empty, error) {
	panic("AuthService.Logout not implemented")
}

// User Management

func (*Server) UpdateContactEmail(ctx context.Context, req *authpb.UpdateContactEmailRequest) (*emptypb.Empty, error) {
	panic("AuthService.UpdateContactEmail not implemented")
}

func (*Server) VerifyContactEmail(ctx context.Context, req *authpb.VerifyContactEmailRequest) (*emptypb.Empty, error) {
	panic("AuthService.VerifyContactEmail not implemented")
}

func (*Server) GetMe(ctx context.Context, req *authpb.GetMeRequest) (*authpb.UserResponse, error) {
	panic("AuthService.GetMe not implemented")
}
