package domain

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Session struct {
	id           uuid.UUID
	username     string
	refreshToken string
	userAgent    string
	clientIp     string
	isBlocked    bool
	expiresAt    time.Time
	createdAt    time.Time
}

func NewSession(username, refreshToken, userAgent, clientIp string, isBlocked bool, expiresAt time.Time) (session *Session, err error) {
	session = &Session{
		id:           uuid.New(),
		username:     username,
		refreshToken: refreshToken,
		userAgent:    userAgent,
		clientIp:     clientIp,
		isBlocked:    isBlocked,
		expiresAt:    expiresAt,
		createdAt:    time.Now(),
	}

	if err = session.validate(); err != nil {
		return
	}

	return
}

func RestoreSession(id uuid.UUID, username, refreshToken, userAgent, clientIp string, isBlocked bool, expiresAt, createdAt time.Time) (session *Session, err error) {
	session = &Session{
		id:           id,
		username:     username,
		refreshToken: refreshToken,
		userAgent:    userAgent,
		clientIp:     clientIp,
		isBlocked:    isBlocked,
		expiresAt:    expiresAt,
		createdAt:    createdAt,
	}

	if err = session.validate(); err != nil {
		return
	}

	return
}

func (s *Session) validate() error {
	if len(s.username) == 0 {
		return errors.New("username is required")
	}
	return nil
}

func (s *Session) ID() uuid.UUID {
	return s.id
}

func (s *Session) Username() string {
	return s.username
}

func (s *Session) RefreshToken() string {
	return s.refreshToken
}

func (s *Session) UserAgent() string {
	return s.userAgent
}

func (s *Session) ClientIp() string {
	return s.clientIp
}

func (s *Session) IsBlocked() bool {
	return s.isBlocked
}

func (s *Session) ExpiresAt() time.Time {
	return s.expiresAt
}

func (s *Session) CreatedAt() time.Time {
	return s.createdAt
}
