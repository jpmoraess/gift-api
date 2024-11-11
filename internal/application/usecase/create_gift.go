package usecase

import (
	"github.com/google/uuid"
	"github.com/jpmoraess/gift-api/internal/application/repository"
	"github.com/jpmoraess/gift-api/internal/domain"
)

type CreateGiftInput struct {
	Gifter    string `json:"gifter"`
	Recipient string `json:"recipient"`
}

type CreateGiftOutput struct {
	ID        uuid.UUID `json:"id"`
	Gifter    string    `json:"gifter"`
	Recipient string    `json:"recipient"`
}

type CreateGift struct {
	giftRepo repository.GiftRepository
}

func NewCreateGift(giftRepo repository.GiftRepository) *CreateGift {
	return &CreateGift{giftRepo: giftRepo}
}

func (uc *CreateGift) Execute(input *CreateGiftInput) (output *CreateGiftOutput, err error) {
	gift, err := domain.NewGift(input.Gifter, input.Recipient)
	if err != nil {
		return
	}

	output = &CreateGiftOutput{
		ID:        gift.ID(),
		Gifter:    gift.Gifter(),
		Recipient: gift.Recipient(),
	}

	return
}
