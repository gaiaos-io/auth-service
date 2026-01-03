package session

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	id                   uuid.UUID
	accountID            uuid.UUID
	refreshToken         RefreshToken
	previousRefreshToken *RefreshToken
	rotatedAt            *time.Time
	expiresAt            time.Time
	revokedAt            *time.Time
	deviceMetadata       DeviceMetadata
	createdAt            time.Time
}

func NewSession(accountID uuid.UUID, refreshToken RefreshToken, deviceMetadata DeviceMetadata, now time.Time, ttl time.Duration) *Session {
	expiresAt := now.Add(ttl)

	return &Session{
		id:                   uuid.New(),
		accountID:            accountID,
		refreshToken:         refreshToken,
		previousRefreshToken: nil,
		rotatedAt:            nil,
		expiresAt:            expiresAt,
		revokedAt:            nil,
		deviceMetadata:       deviceMetadata,
		createdAt:            now,
	}
}

func (session Session) isExpired(now time.Time) bool {
	return now.After(session.expiresAt)
}

func (session Session) isRevoked() bool {
	return session.revokedAt != nil
}

func (session Session) IsActive(now time.Time) bool {
	return !session.isExpired(now) && !session.isRevoked()
}

func (session Session) CanModify(now time.Time) error {
	if !session.IsActive(now) {
		return errors.New("cannot modify session after it's been revoked or has expired")
	}
	return nil
}

func (session *Session) RotateRefreshToken(newRefreshToken RefreshToken, now time.Time) error {
	if err := session.CanModify(now); err != nil {
		return err
	}

	if session.refreshToken.Equal(newRefreshToken) {
		return errors.New("new refresh token can't be equal to curent one")
	}

	session.previousRefreshToken = &session.refreshToken
	session.refreshToken = newRefreshToken
	session.rotatedAt = &now

	return nil
}

func (session *Session) Revoke(now time.Time) error {
	if err := session.CanModify(now); err != nil {
		return err
	}
	session.revokedAt = &now
	return nil
}
