package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jpmoraess/gift-api/internal/domain"
)

type SessionRepository interface {
	Save(ctx context.Context, session *domain.Session) (err error)

	GetSession(ctx context.Context, id uuid.UUID) (session *domain.Session, err error)
}
