package ports

import "github.com/gaiaos-io/auth-service/internal/domain/credential"

type PasswordHasher interface {
	Hash(plainPassword credential.PlainPassword) (*credential.PasswordHash, error)
	Verify(passwordHash credential.PasswordHash, plainPassword credential.PlainPassword) (bool, error)
}
