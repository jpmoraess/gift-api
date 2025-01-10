package chain

import "context"

// GenerateChargeInput - represents object to generate charge
type GenerateChargeInput struct {
	Amount  float64 `json:"amount"`
	DueDate string  `json:"dueDate"`
	// ExternalReference - represents my application transaction id into gateway
	ExternalReference string `json:"externalReference"`
}

// GenerateChargeOutput - represents response object after generate process
type GenerateChargeOutput struct {
	ID string `json:"id"`
}

// ChargeGenerator - defines the GenerateCharge method and SetNext method to indicate the next generator in chain
type ChargeGenerator interface {
	GenerateCharge(ctx context.Context, input *GenerateChargeInput) (output *GenerateChargeOutput, err error)
	SetNext(generator ChargeGenerator)
}

// BaseChargeGenerator - base structure to implement chain pattern
type BaseChargeGenerator struct {
	next ChargeGenerator
}

// SetNext - configure the next processor in chain
func (b *BaseChargeGenerator) SetNext(generator ChargeGenerator) {
	b.next = generator
}
