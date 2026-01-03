package credential

import "errors"

type PasswordHash struct {
	value string
}

func NewPasswordHash(hash string) (*PasswordHash, error) {
	passwordHash := &PasswordHash{value: hash}
	if err := passwordHash.IsValid(); err != nil {
		return nil, err
	}
	return passwordHash, nil
}

func (hash PasswordHash) IsValid() error {
	if hash.value == "" {
		return errors.New("password hash cannot be empty")
	}
	return nil
}

func (hash PasswordHash) String() string {
	return hash.value
}
