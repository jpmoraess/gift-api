package usecase

import (
	"context"
	"fmt"
	"github.com/jpmoraess/gift-api/internal/application/repository"
	"github.com/jpmoraess/gift-api/internal/domain"
	"github.com/jpmoraess/gift-api/token"
	"time"
)

type GenerateTokenInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GenerateTokenOutput struct {
	Token string `json:"token"`
}

type GenerateToken struct {
	tokenMaker     token.Maker
	userRepository repository.UserRepository
}

func NewGenerateToken(tokenMaker token.Maker, userRepository repository.UserRepository) *GenerateToken {
	return &GenerateToken{tokenMaker: tokenMaker, userRepository: userRepository}
}

func (g *GenerateToken) Execute(ctx context.Context, input *GenerateTokenInput) (output *GenerateTokenOutput, err error) {
	user, err := g.userRepository.GetUserByUsername(ctx, input.Username)
	if err != nil {
		fmt.Println("user not found", err.Error())
		return
	}

	err = domain.CheckPassword(input.Password, user.Password())
	if err != nil {
		fmt.Println("user password error", err.Error())
		return
	}

	accessToken, err := g.tokenMaker.CreateToken(input.Username, time.Hour*12)
	if err != nil {
		fmt.Println("error generating access token", err.Error())
		return
	}

	output = &GenerateTokenOutput{
		Token: accessToken,
	}

	return
}
