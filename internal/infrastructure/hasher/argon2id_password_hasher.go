package hasher

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/crypto/argon2"

	"github.com/gaiaos-io/auth-service/internal/application/ports"
	"github.com/gaiaos-io/auth-service/internal/domain/credential"
)

var _ ports.PasswordHasher = (*Argon2idPasswordHasher)(nil)

type Argon2idPasswordHasher struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	keyLength   uint32
	saltLength  uint32
}

func NewArgon2idPasswordHasher(memory, iterations, keyLength, saltLenght uint32, parallelism uint8) (*Argon2idPasswordHasher, error) {
	return &Argon2idPasswordHasher{
		memory:      memory,
		iterations:  iterations,
		parallelism: parallelism,
		keyLength:   keyLength,
		saltLength:  saltLenght,
	}, nil
}

func (hasher Argon2idPasswordHasher) Hash(
	plainPassword credential.PlainPassword,
) (*credential.PasswordHash, error) {
	salt := make([]byte, hasher.saltLength)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}

	hash := argon2.IDKey(
		[]byte(plainPassword.String()),
		salt,
		hasher.iterations,
		hasher.memory,
		hasher.parallelism,
		hasher.keyLength,
	)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	hashedPasswordStr := fmt.Sprintf(
		"$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s",
		hasher.memory,
		hasher.iterations,
		hasher.parallelism,
		b64Salt,
		b64Hash,
	)

	return credential.NewPasswordHash(hashedPasswordStr)
}

var ErrInvalidPasswordHashFormat = errors.New("invalid password hash format")

func (hasher Argon2idPasswordHasher) Verify(
	passwordHash credential.PasswordHash,
	plainPassword credential.PlainPassword,
) (bool, error) {
	parts := strings.Split(passwordHash.String(), "$")
	if len(parts) != 6 {
		return false, ErrInvalidPasswordHashFormat
	}

	// parts:
	// 0 = ""
	// 1 = "argon2id"
	// 2 = "v=19"
	// 3 = "m=65536,t=3,p=2"
	// 4 = salt (b64)
	// 5 = hash (b64)

	if parts[1] != "argon2id" {
		return false, ErrInvalidPasswordHashFormat
	}

	if parts[2] != "v=19" {
		return false, ErrInvalidPasswordHashFormat
	}

	var memory uint32
	var iterations uint32
	var parallelism uint8

	params := strings.Split(parts[3], ",")
	for _, param := range params {
		keyValue := strings.Split(param, "=")
		if len(keyValue) != 2 {
			return false, ErrInvalidPasswordHashFormat
		}

		key := keyValue[0]
		value := keyValue[1]

		switch key {
		case "m":
			memVal, err := strconv.ParseUint(value, 10, 32)
			if err != nil {
				return false, err
			}
			memory = uint32(memVal)
		case "t":
			iterVal, err := strconv.ParseUint(value, 10, 32)
			if err != nil {
				return false, err
			}
			iterations = uint32(iterVal)
		case "p":
			parallelValue, err := strconv.ParseUint(value, 8, 8)
			if err != nil {
				return false, err
			}
			parallelism = uint8(parallelValue)
		default:
			return false, ErrInvalidPasswordHashFormat
		}
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}

	expectedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}

	computedHash := argon2.IDKey(
		[]byte(plainPassword.String()),
		salt,
		iterations,
		memory,
		parallelism,
		uint32(len(expectedHash)),
	)

	if subtle.ConstantTimeCompare(computedHash, expectedHash) == 1 {
		return true, nil
	}

	return false, nil
}
