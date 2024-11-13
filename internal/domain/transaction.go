package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type TransactionStatus string

const (
	TransactionPending   TransactionStatus = "PENDING"
	TransactionPaid      TransactionStatus = "PAID"
	TransactionFailed    TransactionStatus = "FAILED"
	TransactionCancelled TransactionStatus = "CANCELLED"
)

type Transaction struct {
	id     uuid.UUID
	giftID uuid.UUID
	amount float64
	date   time.Time
	status TransactionStatus
}

func NewTransaction(giftID uuid.UUID, amount float64) (transaction *Transaction, err error) {
	transaction = &Transaction{
		id:     uuid.New(),
		giftID: giftID,
		amount: amount,
		date:   time.Now(),
		status: TransactionPending,
	}

	if err = transaction.validate(); err != nil {
		return
	}

	if err = transaction.validateInitialStatus(); err != nil {
		return
	}

	return
}

func (t *Transaction) validate() error {
	if t.id == uuid.Nil {
		return errors.New("id is required")
	}
	if t.giftID == uuid.Nil {
		return errors.New("giftID is required")
	}
	if t.amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	return nil
}

func (t *Transaction) validateInitialStatus() error {
	if t.status != TransactionPending {
		return errors.New("status is not in correct state to initialize")
	}
	return nil
}

func (t *Transaction) Pay() error {
	if t.status != TransactionPending {
		return errors.New("status is not in correct state to pay operation")
	}
	return nil
}

func (t *Transaction) ID() uuid.UUID {
	return t.id
}

func (t *Transaction) GiftID() uuid.UUID {
	return t.giftID
}

func (t *Transaction) Amount() float64 {
	return t.amount
}

func (t *Transaction) Date() time.Time {
	return t.date
}

func (t *Transaction) Status() TransactionStatus {
	return t.status
}
