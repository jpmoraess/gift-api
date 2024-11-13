package repository

import (
	"context"
	"github.com/jpmoraess/gift-api/internal/domain"
)

type TransactionRepository interface {
	Save(ctx context.Context, transaction *domain.Transaction) (err error)
}
