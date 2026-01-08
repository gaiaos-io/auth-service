package hasher

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"errors"

	"github.com/gaiaos-io/auth-service/internal/application/ports"
)

var _ ports.TokenHasher = (*HmacSHA256TokenHasher)(nil)

type HmacSHA256TokenHasher struct {
	secret []byte
}

func NewHmacSHA256TokenHasher(secret []byte) *HmacSHA256TokenHasher {
	s := make([]byte, len(secret))
	copy(s, secret)

	return &HmacSHA256TokenHasher{secret: s}
}

type tokenHash [32]byte

func (hasher HmacSHA256TokenHasher) Hash(token []byte) []byte {
	mac := hmac.New(sha256.New, hasher.secret)
	mac.Write(token)

	var out tokenHash
	copy(out[:], mac.Sum(nil))

	return out[:]
}

func (hasher HmacSHA256TokenHasher) Verify(hash, token []byte) (bool, error) {
	if len(hash) != sha256.Size {
		return false, errors.New("invalid token hash length")
	}

	var expectedHash tokenHash
	copy(expectedHash[:], hash)

	var computedHash tokenHash
	copy(computedHash[:], hasher.Hash(token))

	return subtle.ConstantTimeCompare(computedHash[:], expectedHash[:]) == 1, nil
}
