package credential

import "github.com/google/uuid"

type PasswordCredential struct {
	authIdentityID uuid.UUID
	passwordHash   PasswordHash
}

func NewPasswordCredential(authIdentityID uuid.UUID, passwordHash PasswordHash) (*PasswordCredential, error) {
	if err := passwordHash.IsValid(); err != nil {
		return nil, err
	}
	return &PasswordCredential{authIdentityID: authIdentityID, passwordHash: passwordHash}, nil
}
