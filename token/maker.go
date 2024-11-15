package token

import "time"

// Maker is and interface for managing tokens
type Maker interface {

	// CreateToken - creates a new token for a specific username and duration
	CreateToken(username string, duration time.Duration) (token string, err error)

	// VerifyToken - checks if the token is valid or not
	VerifyToken(token string) (payload *Payload, err error)
}
