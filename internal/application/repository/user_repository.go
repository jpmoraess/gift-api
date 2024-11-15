package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jpmoraess/gift-api/internal/domain"
)

type UserRepository interface {
	Save(ctx context.Context, user *domain.User) (err error)

	GetUser(ctx context.Context, id uuid.UUID) (user *domain.User, err error)

	GetUserByEmail(ctx context.Context, email string) (user *domain.User, err error)

	GetUserByUsername(ctx context.Context, username string) (user *domain.User, err error)
}
