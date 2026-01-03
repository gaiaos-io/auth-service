package identity

import (
	"errors"

	"github.com/gaiaos-io/auth-service/internal/domain/shared"
)

type AuthProvider string

const (
	ProviderEmailPassword AuthProvider = "email_password"
	ProviderGoogle        AuthProvider = "google"
	ProviderGitHub        AuthProvider = "github"
)

func (provider AuthProvider) IsValid() error {
	switch provider {
	case ProviderEmailPassword, ProviderGoogle, ProviderGitHub:
		return nil
	default:
		return errors.New("invalid auth provider")
	}
}

func (provider AuthProvider) ValidateProviderData(providerSubject ProviderSubject, providerEmail *shared.EmailAddress) error {
	if err := providerSubject.IsValid(); err != nil {
		return err
	}

	if providerEmail != nil {
		if err := providerEmail.IsValid(); err != nil {
			return err
		}
	}

	if provider == ProviderEmailPassword &&
		(providerEmail == nil || providerEmail.String() != providerSubject.value) {
		return errors.New("for EmailPassword provider, providerEmail is required and it must have the same value as providerSubject")
	}

	return nil
}

func (provider AuthProvider) IsOAuth() bool {
	return provider != ProviderEmailPassword
}

func (provider AuthProvider) RequiresPassword() bool {
	return provider == ProviderEmailPassword
}
