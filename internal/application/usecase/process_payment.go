package usecase

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jpmoraess/gift-api/internal/application/repository"
	"github.com/jpmoraess/gift-api/internal/domain"
	"github.com/jpmoraess/gift-api/internal/infra/chain"
)

type ProcessPaymentInput struct {
	GiftID uuid.UUID `json:"giftId"`
	Amount float64   `json:"amount"`
}

type ProcessPaymentOutput struct {
	ID     uuid.UUID `json:"id"`
	GiftID uuid.UUID `json:"giftId"`
	Amount float64   `json:"amount"`
}

type ProcessPayment struct {
	paymentProcessor      chain.PaymentProcessor
	transactionRepository repository.TransactionRepository
}

func NewProcessPayment(
	paymentProcessor chain.PaymentProcessor,
	transactionRepository repository.TransactionRepository,
) *ProcessPayment {
	return &ProcessPayment{
		paymentProcessor:      paymentProcessor,
		transactionRepository: transactionRepository,
	}
}

func (p *ProcessPayment) Execute(ctx context.Context, input *ProcessPaymentInput) (output *ProcessPaymentOutput, err error) {
	transaction, err := domain.NewTransaction(input.GiftID, input.Amount)
	if err != nil {
		fmt.Println("error while creating transaction:", err)
		return
	}

	_, err = p.paymentProcessor.ProcessPayment(ctx, &chain.ProcessPaymentInput{
		Amount: transaction.Amount(),
	})
	if err != nil {
		fmt.Println("error while processing payment:", err)
		return
	}

	err = p.transactionRepository.Save(ctx, transaction)
	if err != nil {
		fmt.Println("error while saving transaction:", err)
		return
	}

	output = &ProcessPaymentOutput{
		ID:     transaction.ID(),
		GiftID: transaction.GiftID(),
		Amount: transaction.Amount(),
	}

	return
}
