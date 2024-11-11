package infra

import (
	"context"
	db "github.com/jpmoraess/gift-api/db/sqlc"

	"github.com/google/uuid"
	"github.com/jpmoraess/gift-api/internal/domain"
)

type GiftRepositoryAdapter struct {
	store db.Store
}

func NewGiftRepositoryAdapter(store db.Store) *GiftRepositoryAdapter {
	return &GiftRepositoryAdapter{store: store}
}

func (g *GiftRepositoryAdapter) Save(ctx context.Context, gift *domain.Gift) (err error) {
	arg := db.InsertGiftParams{
		ID:        gift.ID(),
		Gifter:    gift.Gifter(),
		Recipient: gift.Recipient(),
		Message:   gift.Message(),
		Status:    db.GiftStatusGiftPending,
	}

	_, err = g.store.InsertGift(ctx, arg)
	if err != nil {
		return err
	}

	return
}

func (g *GiftRepositoryAdapter) Get(ctx context.Context, id uuid.UUID) (gift *domain.Gift, err error) {
	return
}
