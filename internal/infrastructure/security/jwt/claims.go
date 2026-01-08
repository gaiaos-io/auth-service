package jwt

import "github.com/gaiaos-io/auth-service/internal/domain/account"

type jwtClaims struct {
	AccountID string                `json:"sub"`
	Roles     []account.AccountRole `json:"roles"`

	Issuer   string `json:"iss"`
	Audience string `json:"aud"`

	IssuedAt  int64 `json:"iat"`
	ExpiresAt int64 `json:"exp"`
}
