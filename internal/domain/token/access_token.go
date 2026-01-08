package token

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/gaiaos-io/auth-service/internal/domain/account"
)

var ErrInvalidAccessToken = errors.New("invalid access token")
var ErrInvalidAccessTokenClaims = errors.New("invalid access token claims")

type AccessTokenClaims struct {
	accountID uuid.UUID
	roles     []account.AccountRole
	expiresAt time.Time
}

func NewAccessTokenClaims(
	accountID uuid.UUID,
	roles []account.AccountRole,
	expiresAt time.Time,
) (AccessTokenClaims, error) {
	if accountID == uuid.Nil {
		return AccessTokenClaims{}, ErrInvalidAccessTokenClaims
	}
	if len(roles) == 0 {
		return AccessTokenClaims{}, ErrInvalidAccessTokenClaims
	}

	return AccessTokenClaims{
		accountID: accountID,
		roles:     roles,
		expiresAt: expiresAt,
	}, nil
}

func (claims AccessTokenClaims) AccountID() uuid.UUID {
	return claims.accountID
}

func (claims AccessTokenClaims) Roles() []account.AccountRole {
	return claims.roles
}

func (claims AccessTokenClaims) ExpiresAt() time.Time {
	return claims.expiresAt
}

type AccessToken struct {
	value string
}

func NewAccessToken(value string) (*AccessToken, error) {
	if value == "" {
		return nil, ErrInvalidAccessToken
	}
	return &AccessToken{value: value}, nil
}

func (token AccessToken) String() string {
	return token.value
}
