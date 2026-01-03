package account

import (
	"slices"

	"github.com/google/uuid"

	"github.com/gaiaos-io/auth-service/internal/domain/shared"
)

type Account struct {
	id           uuid.UUID
	status       AccountStatus
	roles        []AccountRole
	ContactEmail *shared.EmailAddress
}

func NewAccount(status AccountStatus) (*Account, error) {
	if err := status.IsValid(); err != nil {
		return nil, err
	}
	if err := status.IsValidInitial(); err != nil {
		return nil, err
	}

	return &Account{
		id:     uuid.New(),
		status: status,
		roles:  []AccountRole{AccountRoleCitizen},
	}, nil
}

// Getters

func (account *Account) ID() uuid.UUID {
	return account.id
}

func (account *Account) Status() AccountStatus {
	return account.status
}

// Roles

func (account *Account) AddRole(newRole AccountRole) error {
	if err := newRole.IsValid(); err != nil {
		return err
	}

	if !slices.Contains(account.roles, newRole) {
		account.roles = append(account.roles, newRole)
	}

	return nil
}

func (account *Account) RemoveRole(roleToBeRemoved AccountRole) error {
	if err := roleToBeRemoved.IsValid(); err != nil {
		return err
	}

	index := slices.Index(account.roles, roleToBeRemoved)

	if index != -1 {
		account.roles = slices.Delete(account.roles, index, index+1)
	}

	return nil
}

func (account Account) HasRole(role AccountRole) (bool, error) {
	if err := role.IsValid(); err != nil {
		return false, err
	}

	return slices.Contains(account.roles, role), nil
}

// Status

func (account *Account) Activate() {
	account.status = AccountStatusActive
}

func (account *Account) Disable() {
	account.status = AccountStatusDisabled
}
