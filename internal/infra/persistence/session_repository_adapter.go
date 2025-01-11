package persistence

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	db "github.com/jpmoraess/gift-api/db/sqlc"
	"github.com/jpmoraess/gift-api/internal/domain"
)

type SessionRepositoryAdapter struct {
	store db.Store
}

func NewSessionRepositoryAdapter(store db.Store) *SessionRepositoryAdapter {
	return &SessionRepositoryAdapter{store: store}
}

func (s *SessionRepositoryAdapter) Save(ctx context.Context, session *domain.Session) (err error) {
	arg := db.CreateSessionParams{
		ID:           session.ID(),
		Username:     session.Username(),
		RefreshToken: session.RefreshToken(),
		UserAgent:    session.UserAgent(),
		ClientIp:     session.ClientIp(),
		IsBlocked:    session.IsBlocked(),
		ExpiresAt:    session.ExpiresAt(),
	}

	_, err = s.store.CreateSession(ctx, arg)
	if err != nil {
		fmt.Println("failed to save session into database")
		return
	}

	return
}

func (s *SessionRepositoryAdapter) GetSession(ctx context.Context, id uuid.UUID) (session *domain.Session, err error) {
	data, err := s.store.GetSession(ctx, id)
	if err != nil {
		fmt.Println("failed to fetch session from database")
		return
	}

	session, err = domain.RestoreSession(
		data.ID,
		data.Username,
		data.RefreshToken,
		data.UserAgent,
		data.ClientIp,
		data.IsBlocked,
		data.ExpiresAt,
		data.CreatedAt,
	)
	if err != nil {
		fmt.Println("failed to restore session from database to domain", err)
		return
	}

	return
}
