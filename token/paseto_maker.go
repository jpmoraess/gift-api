package token

import (
	"fmt"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
	"time"
)

// PasetMaker is a PASETO token maker
type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

// NewPasetomaker creates a PasetoMaker
func NewPasetoMaker(symmetricKey []byte) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("symmetricKey must be 32 bytes long")
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return maker, nil
}

func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (token string, err error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return
	}

	return maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
}

func (maker *PasetoMaker) VerifyToken(token string) (payload *Payload, err error) {
	payload = &Payload{}

	err = maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return
	}

	err = payload.Valid()
	if err != nil {
		return
	}

	return
}
