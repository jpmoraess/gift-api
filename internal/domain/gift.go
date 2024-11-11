package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Gift struct {
	id        uuid.UUID
	gifter    string
	recipient string
	createdAt time.Time
}

func NewGift(gifter, recipient string) (gift *Gift, err error) {
	gift = &Gift{
		id:        uuid.New(),
		gifter:    gifter,
		recipient: recipient,
		createdAt: time.Now(),
	}

	if err = gift.validate(); err != nil {
		return
	}

	return
}

func (g *Gift) validate() error {
	if g.id == uuid.Nil {
		return errors.New("id is required")
	}
	if len(g.gifter) == 0 {
		return errors.New("gifter is required")
	}
	if len(g.recipient) == 0 {
		return errors.New("recipient is required")
	}
	return nil
}

func (g *Gift) ID() uuid.UUID {
	return g.id
}

func (g *Gift) Gifter() string {
	return g.gifter
}

func (g *Gift) Recipient() string {
	return g.recipient
}

func (g *Gift) CreatedAt() time.Time {
	return g.createdAt
}
