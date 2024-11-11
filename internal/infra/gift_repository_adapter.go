package infra

import (
	"context"

	"github.com/google/uuid"
	"github.com/jpmoraess/gift-api/internal/domain"
)

type GiftRepositoryAdapter struct {
}

func NewGiftRepositoryAdapter() *GiftRepositoryAdapter {
	return &GiftRepositoryAdapter{}
}

func (g *GiftRepositoryAdapter) Save(ctx context.Context, gift *domain.Gift) (err error) {
	return
}

func (g *GiftRepositoryAdapter) Get(ctx context.Context, id uuid.UUID) (gift *domain.Gift, err error) {
	return
}
