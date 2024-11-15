package persistence

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	db "github.com/jpmoraess/gift-api/db/sqlc"
	"github.com/jpmoraess/gift-api/internal/domain"
)

type UserRepositoryAdapter struct {
	store db.Store
}

func NewUserRepositoryAdapter(store db.Store) *UserRepositoryAdapter {
	return &UserRepositoryAdapter{store: store}
}

func (u *UserRepositoryAdapter) Save(ctx context.Context, user *domain.User) (err error) {
	arg := db.CreateUserParams{
		ID:        user.ID(),
		Username:  user.Username(),
		Password:  user.Password(),
		FullName:  user.FullName(),
		Email:     user.Email(),
		CreatedAt: user.CreatedAt(),
	}

	_, err = u.store.CreateUser(ctx, arg)
	if err != nil {
		fmt.Println("failed to save user to db", err)
		return
	}

	return err
}

func (u *UserRepositoryAdapter) GetUser(ctx context.Context, id uuid.UUID) (user *domain.User, err error) {
	data, err := u.store.GetUser(ctx, id)
	if err != nil {
		fmt.Println("failed to retrieve user from db using id", err)
		return
	}

	user, err = domain.RestoreUser(data.ID, data.Username, data.Password, data.FullName, data.Email, data.CreatedAt)
	if err != nil {
		fmt.Println("failed to restore user db to domain", err)
		return
	}

	return
}

func (u *UserRepositoryAdapter) GetUserByEmail(ctx context.Context, email string) (user *domain.User, err error) {
	data, err := u.store.GetUserByEmail(ctx, email)
	if err != nil {
		fmt.Println("failed to retrieve user from db using email", err)
		return
	}

	user, err = domain.RestoreUser(data.ID, data.Username, data.Password, data.FullName, data.Email, data.CreatedAt)
	if err != nil {
		fmt.Println("failed to restore user db to domain", err)
		return
	}

	return
}

func (u *UserRepositoryAdapter) GetUserByUsername(ctx context.Context, username string) (user *domain.User, err error) {
	data, err := u.store.GetUserByEmail(ctx, username)
	if err != nil {
		fmt.Println("failed to retrieve user from db using username", err)
		return
	}

	user, err = domain.RestoreUser(data.ID, data.Username, data.Password, data.FullName, data.Email, data.CreatedAt)
	if err != nil {
		fmt.Println("failed to restore user db to domain", err)
		return
	}

	return
}
