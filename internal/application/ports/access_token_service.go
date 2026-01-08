package ports

import (
	"time"

	"github.com/gaiaos-io/auth-service/internal/domain/token"
)

type AccessTokenService interface {
	Sign(claims token.AccessTokenClaims, at time.Time) (*token.AccessToken, error)
	Verify(accessToken token.AccessToken) (*token.AccessTokenClaims, error)
}
