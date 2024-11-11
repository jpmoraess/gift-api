package domain

import (
	"errors"

	"github.com/google/uuid"
)

type PaymentStatus int

const (
	PaymentCompleted = iota + 1
	PaymentCancelled
	PaymentFailed
)

type Payment struct {
	id     uuid.UUID
	giftID uuid.UUID
	status PaymentStatus
}

func NewPayment(giftID uuid.UUID) (payment *Payment, err error) {
	payment = &Payment{
		id:     uuid.New(),
		giftID: giftID,
		status: PaymentCompleted,
	}

	if err = payment.validate(); err != nil {
		return
	}

	return
}

func (p *Payment) validate() error {
	if p.id == uuid.Nil {
		return errors.New("id is required")
	}
	if p.giftID == uuid.Nil {
		return errors.New("giftID is required")
	}
	return nil
}

func (p *Payment) ID() uuid.UUID {
	return p.id
}

func (p *Payment) GiftID() uuid.UUID {
	return p.giftID
}
