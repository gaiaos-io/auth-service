package jwt

import (
	"crypto/ecdsa"
	"errors"
	"time"

	jwtlib "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/gaiaos-io/auth-service/internal/application/ports"
	"github.com/gaiaos-io/auth-service/internal/domain/account"
	"github.com/gaiaos-io/auth-service/internal/domain/token"
)

var _ ports.AccessTokenService = (*JwtService)(nil)

type JwtService struct {
	issuer    string
	audience  string
	ttl       time.Duration
	private   *ecdsa.PrivateKey
	public    *ecdsa.PublicKey
	algorithm jwtlib.SigningMethod
}

func NewJwtService(
	issuer, audience string,
	ttl time.Duration,
	privBytes, pubBytes []byte,
) (*JwtService, error) {
	priv, err := parseECDSAPrivateKey(privBytes)
	if err != nil {
		return nil, err
	}

	pub, err := parseECDSAPublicKey(pubBytes)
	if err != nil {
		return nil, err
	}

	return &JwtService{
		issuer:    issuer,
		audience:  audience,
		ttl:       ttl,
		private:   priv,
		public:    pub,
		algorithm: jwtlib.SigningMethodES256,
	}, nil
}

func (service JwtService) Sign(claims token.AccessTokenClaims, at time.Time) (*token.AccessToken, error) {
	jwtClaims := jwtClaims{
		AccountID: claims.AccountID().String(),
		Roles:     claims.Roles(),
		Issuer:    service.issuer,
		Audience:  service.audience,
		IssuedAt:  at.Unix(),
		ExpiresAt: claims.ExpiresAt().Unix(),
	}

	jwtToken := jwtlib.NewWithClaims(service.algorithm, jwtlib.MapClaims{
		"sub":   jwtClaims.AccountID,
		"roles": jwtClaims.Roles,
		"iss":   jwtClaims.Issuer,
		"aud":   jwtClaims.Audience,
		"iat":   jwtClaims.IssuedAt,
		"exp":   jwtClaims.ExpiresAt,
	})

	signed, err := jwtToken.SignedString(service.private)
	if err != nil {
		return nil, err
	}

	return token.NewAccessToken(signed)
}

func (service JwtService) Verify(accessToken token.AccessToken) (*token.AccessTokenClaims, error) {
	parsed, err := jwtlib.Parse(accessToken.String(), func(token *jwtlib.Token) (any, error) {
		if token.Method.Alg() != service.algorithm.Alg() {
			return nil, errors.New("unexpected signing method")
		}
		return service.public, nil
	})
	if err != nil {
		return nil, err
	}
	if !parsed.Valid {
		return nil, errors.New("invalid access token")
	}

	claims, ok := parsed.Claims.(jwtlib.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	accountID, err := uuid.Parse(claims["sub"].(string))
	if err != nil {
		return nil, err
	}

	roles := claims["roles"].([]account.AccountRole)
	for _, role := range roles {
		if err := role.IsValid(); err != nil {
			return nil, err
		}
	}

	exp := time.Unix(int64(claims["exp"].(float64)), 0)

	tokenClaims, err := token.NewAccessTokenClaims(accountID, roles, exp)
	if err != nil {
		return nil, err
	}

	return &tokenClaims, nil
}
