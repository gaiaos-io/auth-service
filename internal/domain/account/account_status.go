package account

import "errors"

type AccountStatus string

const (
	AccountStatusUnverified AccountStatus = "unverified"
	AccountStatusActive     AccountStatus = "active"
	AccountStatusDisabled   AccountStatus = "disabled"
)

func (status AccountStatus) IsValid() error {
	switch status {
	case AccountStatusUnverified, AccountStatusActive, AccountStatusDisabled:
		return nil
	default:
		return errors.New("invalid account status")
	}
}

func (status AccountStatus) IsValidInitial() error {
	if status != AccountStatusActive && status != AccountStatusUnverified {
		return errors.New("initial account status may only be Active (OAuth) or Unverified (Email-Password)")
	}
	return nil
}

func (status AccountStatus) CanLogin() error {
	if status != AccountStatusActive {
		return errors.New("only Active account status may login")
	}
	return nil
}
