package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jpmoraess/gift-api/internal/domain"
)

type GiftRepository interface {
	Save(ctx context.Context, gift *domain.Gift) (err error)
	Get(ctx context.Context, id uuid.UUID) (gift *domain.Gift, err error)
}
