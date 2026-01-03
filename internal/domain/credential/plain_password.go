package credential

import (
	"errors"
	"regexp"
)

var validPasswordPattern = regexp.MustCompile(
	`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`,
)

type PlainPassword struct {
	value string
}

func NewPlainPassword(password string) (*PlainPassword, error) {
	plainPassword := &PlainPassword{value: password}
	if err := plainPassword.IsValid(); err != nil {
		return nil, err
	}
	return plainPassword, nil
}

func (password PlainPassword) IsValid() error {
	if !validPasswordPattern.MatchString(password.value) {
		return errors.New("password must be at least 8 characters long, contain an uppercase and a lowercase character, a number, and a symbol")
	}

	return nil
}

// String returns the raw password.
// Must only be used for hashing and never persisted or logged.
func (password PlainPassword) String() string {
	return password.value
}
