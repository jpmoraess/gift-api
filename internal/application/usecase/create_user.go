package usecase

import (
	"context"
	"fmt"
	"github.com/jpmoraess/gift-api/internal/application/repository"
	"github.com/jpmoraess/gift-api/internal/domain"
)

type CreateUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
}

type CreateUserOutput struct {
	Username string `json:"username"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
}

type CreateUser struct {
	userRepository repository.UserRepository
}

func NewCreateUser(userRepository repository.UserRepository) *CreateUser {
	return &CreateUser{userRepository: userRepository}
}

func (uc *CreateUser) Execute(ctx context.Context, input *CreateUserInput) (output *CreateUserOutput, err error) {
	password, err := domain.HashPassword(input.Password)
	if err != nil {
		fmt.Println("error while hashing password:", err.Error())
		return
	}

	user, err := domain.NewUser(input.Username, password, input.FullName, input.Email)
	if err != nil {
		fmt.Println("error while creating user:", err.Error())
		return
	}

	err = uc.userRepository.Save(ctx, user)
	if err != nil {
		fmt.Println("error while saving user:", err.Error())
		return
	}

	output = &CreateUserOutput{
		Username: user.Username(),
		FullName: user.FullName(),
		Email:    user.Email(),
	}

	return
}
