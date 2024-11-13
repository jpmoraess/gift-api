package usecase

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jpmoraess/gift-api/internal/application/repository"
	"github.com/jpmoraess/gift-api/internal/domain"
)

type CreateGiftInput struct {
	Gifter    string `json:"gifter"`
	Recipient string `json:"recipient"`
	Message   string `json:"message"`
}

type CreateGiftOutput struct {
	ID        uuid.UUID `json:"id"`
	Gifter    string    `json:"gifter"`
	Recipient string    `json:"recipient"`
}

type CreateGift struct {
	giftRepository repository.GiftRepository
}

func NewCreateGift(giftRepository repository.GiftRepository) *CreateGift {
	return &CreateGift{giftRepository: giftRepository}
}

func (uc *CreateGift) Execute(ctx context.Context, input *CreateGiftInput) (output *CreateGiftOutput, err error) {
	gift, err := domain.NewGift(input.Gifter, input.Recipient, input.Message)
	if err != nil {
		fmt.Println("error while creating gift:", err)
		return
	}

	err = uc.giftRepository.Save(ctx, gift)
	if err != nil {
		fmt.Println("error while saving gift:", err)
		return
	}

	output = &CreateGiftOutput{
		ID:        gift.ID(),
		Gifter:    gift.Gifter(),
		Recipient: gift.Recipient(),
	}

	return
}
