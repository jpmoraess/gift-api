package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/jpmoraess/gift-api/internal/application/repository"
	"github.com/jpmoraess/gift-api/internal/domain"
	"github.com/jpmoraess/gift-api/internal/infra/chain"
)

// GenerateChargeInput represents the input for creating a transaction
// @Description GenerateChargeInput represents the input for creating a transaction
// @Model
type GenerateChargeInput struct {
	Amount float64 `json:"amount"`
}

type GenerateChargeOutput struct {
	ID     uuid.UUID `json:"id"`
	Amount float64   `json:"amount"`
}

type GenerateCharge struct {
	generator             chain.ChargeGenerator
	transactionRepository repository.TransactionRepository
}

func NewGenerateCharge(
	generator chain.ChargeGenerator,
	transactionRepository repository.TransactionRepository,
) *GenerateCharge {
	return &GenerateCharge{
		generator:             generator,
		transactionRepository: transactionRepository,
	}
}

func (g *GenerateCharge) Execute(ctx context.Context, input *GenerateChargeInput) (output *GenerateChargeOutput, err error) {
	transaction, err := domain.NewTransaction(input.Amount)
	if err != nil {
		return
	}

	charge, err := g.generator.GenerateCharge(ctx, &chain.GenerateChargeInput{
		Amount:            transaction.Amount(),
		DueDate:           transaction.DueDate().Format("2006-01-02"),
		ExternalReference: transaction.ID().String(),
	})
	if err != nil {
		return
	}
	transaction.SetExternalID(charge.ID)

	err = g.transactionRepository.Save(ctx, transaction)
	if err != nil {
		return
	}

	output = &GenerateChargeOutput{
		ID:     transaction.ID(),
		Amount: transaction.Amount(),
	}

	return
}
