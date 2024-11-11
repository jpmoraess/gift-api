package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type GiftStatus int

const (
	GIFT_PENDING GiftStatus = iota + 1
	GIFT_PAID
	GIFT_APPROVED
	GIFT_CANCELLING
	GIFT_CANCELLED
)

type Gift struct {
	id        uuid.UUID
	gifter    string
	recipient string
	message   string
	status    GiftStatus
	createdAt time.Time
}

func NewGift(gifter, recipient, message string) (gift *Gift, err error) {
	gift = &Gift{
		id:        uuid.New(),
		gifter:    gifter,
		recipient: recipient,
		message:   message,
		status:    GIFT_PENDING,
		createdAt: time.Now(),
	}

	if err = gift.validate(); err != nil {
		return
	}

	if err = gift.validateInitialStatus(); err != nil {
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
	if len(g.message) == 0 {
		return errors.New("message is required")
	}
	return nil
}

func (g *Gift) validateInitialStatus() error {
	if g.status != GIFT_PENDING {
		return errors.New("gift is not in correct state to initialize")
	}
	return nil
}

func (g *Gift) Pay() error {
	if g.status != GIFT_PENDING {
		return errors.New("gift is not in correct state for pay operation")
	}
	g.status = GIFT_PAID
	return nil
}

func (g *Gift) Approve() error {
	if g.status != GIFT_PAID {
		return errors.New("gift is not in correct state for approve operation")
	}
	g.status = GIFT_APPROVED
	return nil
}

func (g *Gift) InitCancel() error {
	if g.status != GIFT_PAID {
		return errors.New("gift is not in correct state for initCancel operation")
	}
	g.status = GIFT_CANCELLING
	return nil
}

func (g *Gift) Cancel() error {
	if !(g.status == GIFT_CANCELLING || g.status == GIFT_PENDING) {
		return errors.New("gift is not in correct state for cancel operation")
	}
	g.status = GIFT_CANCELLED
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

func (g *Gift) Message() string {
	return g.message
}

func (g *Gift) CreatedAt() time.Time {
	return g.createdAt
}
