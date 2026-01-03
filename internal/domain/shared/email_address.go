package shared

import (
	"errors"
	"regexp"
	"strings"
)

var validEmailPattern = regexp.MustCompile(
	`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$`,
)

type EmailAddress struct {
	value string
}

func NewEmailAddress(address string) (*EmailAddress, error) {
	address = strings.TrimSpace(strings.ToLower(address))
	emailAddress := &EmailAddress{value: address}

	if err := emailAddress.IsValid(); err != nil {
		return nil, err
	}

	return emailAddress, nil
}

func (address EmailAddress) IsValid() error {
	if validEmailPattern.MatchString(address.value) {
		return errors.New("invalid email address")
	}
	return nil
}

func (address EmailAddress) String() string {
	return address.value
}

func (address EmailAddress) Equal(other EmailAddress) bool {
	return address.value == other.value
}
