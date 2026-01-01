package grpcserver

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	authpb "github.com/gaiaos-io/auth-service/proto/v1"
)

var _ authpb.AuthServiceServer = (*AuthServiceServer)(nil)

type AuthServiceServer struct {
	authpb.UnimplementedAuthServiceServer
}

func NewAuthServiceServer() *AuthServiceServer {
	return &AuthServiceServer{}
}

// ─────────────────────────────────
// Identity
// ─────────────────────────────────
func (*Server) GetMe(ctx context.Context, req *authpb.GetMeRequest) (*authpb.UserResponse, error) {
	panic("GetMe not implemented")
}

// Email-password
func (*Server) EmailPasswordRegister(ctx context.Context, req *authpb.EmailPasswordRegisterRequest) (*authpb.AuthResponse, error) {
	panic("EmailPasswordRegister not implemented")
}

func (*Server) EmailPasswordLogin(ctx context.Context, req *authpb.EmailPasswordLoginRequest) (*authpb.AuthResponse, error) {
	panic("EmailPasswordLogin not implemented")
}

// OAuth
func (*Server) StartOAuthLogin(ctx context.Context, req *authpb.StartOAuthRequest) (*authpb.StartOAuthResponse, error) {
	panic("StartOAuthLogin not implemented")
}

func (*Server) FinishOAuthLogin(ctx context.Context, req *authpb.FinishOAuthRequest) (*authpb.AuthResponse, error) {
	panic("FinishOAuthLogin not implemented")
}

// Email-password
func (*Server) EmailPasswordReauthenticate(ctx context.Context, req *authpb.EmailPasswordReauthenticateRequest) (*authpb.ReauthenticateResponse, error) {
	panic("EmailPasswordReauthenticate not implemented")
}

// OAuth
func (*Server) StartOAuthReauthenticate(ctx context.Context, req *authpb.StartOAuthRequest) (*authpb.StartOAuthResponse, error) {
	panic("StartOAuthReauthenticate not implemented")
}

func (*Server) FinishOAuthReauthenticate(ctx context.Context, req *authpb.FinishOAuthRequest) (*authpb.ReauthenticateResponse, error) {
	panic("FinishOAuthReauthenticate not implemented")
}

// ─────────────────────────────────
// Session & Token Management
// ─────────────────────────────────
func (*Server) RefreshToken(ctx context.Context, req *authpb.RefreshTokenRequest) (*authpb.AuthResponse, error) {
	panic("RefreshToken not implemented")
}

func (*Server) Logout(ctx context.Context, req *authpb.LogoutRequest) (*emptypb.Empty, error) {
	panic("Logout not implemented")
}

func (*Server) RevokeSession(ctx context.Context, req *authpb.RevokeSessionRequest) (*emptypb.Empty, error) {
	panic("RevokeSession not implemented")
}

func (*Server) ListSessions(ctx context.Context, req *authpb.ListSessionsRequest) (*authpb.ListSessionsResponse, error) {
	panic("ListSessions not implemented")
}

// Password lifecycle
func (*Server) ChangePassword(ctx context.Context, req *authpb.ChangePasswordRequest) (*emptypb.Empty, error) {
	panic("ChangePassword not implemented")
}

func (*Server) RequestPasswordReset(ctx context.Context, req *authpb.RequestPasswordResetRequest) (*emptypb.Empty, error) {
	panic("RequestPasswordReset not implemented")
}

func (*Server) ConfirmPasswordReset(ctx context.Context, req *authpb.ConfirmPasswordResetRequest) (*emptypb.Empty, error) {
	panic("ConfirmPasswordReset not implemented")
}

// Email-password linking
func (*Server) EmailPasswordLink(ctx context.Context, req *authpb.EmailPasswordLinkRequest) (*emptypb.Empty, error) {
	panic("EmailPasswordLink not implemented")
}

// OAuth linking
func (*Server) StartOAuthLink(ctx context.Context, req *authpb.StartOAuthRequest) (*authpb.StartOAuthResponse, error) {
	panic("StartOAuthLink not implemented")
}

func (*Server) FinishOAuthLink(ctx context.Context, req *authpb.FinishOAuthRequest) (*emptypb.Empty, error) {
	panic("FinishOAuthLink not implemented")
}

func (*Server) UnlinkOAuth(ctx context.Context, req *authpb.UnlinkOAuthRequest) (*emptypb.Empty, error) {
	panic("UnlinkOAuth not implemented")
}

// ─────────────────────────────────
// Account Metadata & Verification
// ─────────────────────────────────
func (*Server) RequestContactEmailUpdate(ctx context.Context, req *authpb.RequestContactEmailUpdateRequest) (*emptypb.Empty, error) {
	panic("RequestContactEmailUpdate not implemented")
}

func (*Server) VerifyEmail(ctx context.Context, req *authpb.VerifyEmailRequest) (*emptypb.Empty, error) {
	panic("VerifyEmail not implemented")
}

// ─────────────────────────────────
// Account State & Lifecycle
// ─────────────────────────────────
func (*Server) DisableAccount(ctx context.Context, req *authpb.DisableAccountRequest) (*emptypb.Empty, error) {
	panic("DisableAccount not implemented")
}

func (*Server) ReactivateAccount(ctx context.Context, req *authpb.ReactivateAccountRequest) (*emptypb.Empty, error) {
	panic("ReactivateAccount not implemented")
}

func (*Server) RequestAccountDeletion(ctx context.Context, req *authpb.RequestAccountDeletionRequest) (*emptypb.Empty, error) {
	panic("RequestAccountDeletion not implemented")
}

func (*Server) CancelAccountDeletion(ctx context.Context, req *authpb.CancelAccountDeletionRequest) (*emptypb.Empty, error) {
	panic("CancelAccountDeletion not implemented")
}

// ─────────────────────────────────
// Token Introspection
// ─────────────────────────────────
func (*Server) IntrospectToken(ctx context.Context, req *authpb.IntrospectTokenRequest) (*authpb.IntrospectTokenResponse, error) {
	panic("IntrospectToken not implemented")
}
