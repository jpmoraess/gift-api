package token

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	ErrTokenExpired = errors.New("token expired")
)

// Payload - contains the payload data of the token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// NewPayload - creates a new token payload with a specific username and duration
func NewPayload(username string, duration time.Duration) (payload *Payload, err error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return
	}
	payload = &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return
}

// Valid - checks if the token payload is valid or not
func (payload *Payload) Valid() (err error) {
	if time.Now().After(payload.ExpiredAt) {
		return ErrTokenExpired
	}
	return
}
