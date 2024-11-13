package chain

import "context"

// ProcessPaymentInput - represents object to process payment
type ProcessPaymentInput struct {
	Amount float64 `json:"amount"`
	// ExternalReference - represents my application transaction id into gateway
	ExternalReference string `json:"externalReference"`
}

// ProcessPaymentOutput - represents response object after payment process
type ProcessPaymentOutput struct {
	ID string `json:"id"`
}

// PaymentProcessor - defines the processPayment method and setNext method to indicate the next processor in chain
type PaymentProcessor interface {
	ProcessPayment(ctx context.Context, input *ProcessPaymentInput) (output *ProcessPaymentOutput, err error)
	SetNext(processor PaymentProcessor)
}

// BasePaymentProcessor - base structure to implement chain pattern
type BasePaymentProcessor struct {
	next PaymentProcessor
}

// SetNext - configure the next processor in chain
func (b *BasePaymentProcessor) SetNext(processor PaymentProcessor) {
	b.next = processor
}
