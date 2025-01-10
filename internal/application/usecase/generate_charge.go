package usecase

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jpmoraess/gift-api/internal/application/repository"
	"github.com/jpmoraess/gift-api/internal/domain"
	"github.com/jpmoraess/gift-api/internal/infra/chain"
)

// GenerateChargeInput represents the input for creating a transaction
// @Description ProcessPaymentInput represents the input for creating a transaction
// @Model
type GenerateChargeInput struct {
	GiftID uuid.UUID `json:"giftId"`
	Amount float64   `json:"amount"`
}

type GenerateChargeOutput struct {
	ID     uuid.UUID `json:"id"`
	GiftID uuid.UUID `json:"giftId"`
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
	transaction, err := domain.NewTransaction(input.GiftID, input.Amount)
	if err != nil {
		fmt.Println("error while creating transaction:", err)
		return
	}

	processPaymentOutput, err := g.generator.GenerateCharge(ctx, &chain.GenerateChargeInput{
		Amount:            transaction.Amount(),
		ExternalReference: transaction.ID().String(),
	})
	if err != nil {
		fmt.Println("error while processing payment:", err)
		return
	}
	transaction.SetExternalID(processPaymentOutput.ID)

	err = g.transactionRepository.Save(ctx, transaction)
	if err != nil {
		fmt.Println("error while saving transaction:", err)
		return
	}

	output = &GenerateChargeOutput{
		ID:     transaction.ID(),
		GiftID: transaction.GiftID(),
		Amount: transaction.Amount(),
	}

	return
}
