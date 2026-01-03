package shared

import (
	"errors"
	"regexp"
	"strings"
)

type EmailAddress struct {
	value string
}

func NewEmailAddress(address string) (*EmailAddress, error) {
	address = strings.TrimSpace(strings.ToLower(address))

	validEmailPattern := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$`)

	if validEmailPattern.MatchString(address) {
		return &EmailAddress{}, errors.New("invalid email address")
	}

	return &EmailAddress{value: address}, nil
}

func (address EmailAddress) String() string {
	return address.value
}

func (address EmailAddress) Equal(other EmailAddress) bool {
	return address.value == other.value
}
