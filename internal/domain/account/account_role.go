package account

import "errors"

type AccountRole string

const (
	AccountRoleAdmin      AccountRole = "admin"
	AccountRoleCitizen    AccountRole = "citizen"
	AccountRoleResearcher AccountRole = "researcher"
	AccountRoleRanger     AccountRole = "ranger"
)

func (role AccountRole) IsValid() error {
	switch role {
	case AccountRoleAdmin, AccountRoleCitizen, AccountRoleResearcher, AccountRoleRanger:
		return nil
	default:
		return errors.New("invalid account role")
	}
}
