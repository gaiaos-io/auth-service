package identity

import (
	"time"

	"github.com/google/uuid"

	"github.com/gaiaos-io/auth-service/internal/domain/shared"
)

type AuthIdentity struct {
	id              uuid.UUID
	accountID       uuid.UUID
	provider        AuthProvider
	providerSubject ProviderSubject
	providerEmail   *shared.EmailAddress
	verifiedAt      *time.Time
}

func NewAuthIdentity(accountId uuid.UUID, provider AuthProvider, providerSubject ProviderSubject, providerEmail *shared.EmailAddress, now time.Time) (*AuthIdentity, error) {
	if err := provider.IsValid(); err != nil {
		return nil, err
	}

	if err := provider.ValidateProviderData(providerSubject, providerEmail); err != nil {
		return nil, err
	}

	var verifiedAt *time.Time = nil
	if provider.IsOAuth() {
		verifiedAt = &now
	}

	return &AuthIdentity{
		id:              uuid.New(),
		accountID:       accountId,
		provider:        provider,
		providerSubject: providerSubject,
		providerEmail:   providerEmail,
		verifiedAt:      verifiedAt,
	}, nil
}

func (identity AuthIdentity) Matches(provider AuthProvider, providerSubject ProviderSubject) bool {
	return identity.provider == provider && identity.providerSubject == providerSubject
}
