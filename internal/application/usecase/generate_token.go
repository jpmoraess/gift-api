package usecase

import (
	"context"
	"fmt"
	"github.com/jpmoraess/gift-api/config"
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
	SessionID             string    `json:"session_id"`
	AccessToken           string    `json:"access_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshToken          string    `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
}

type GenerateToken struct {
	config            *config.Config
	tokenMaker        token.Maker
	userRepository    repository.UserRepository
	sessionRepository repository.SessionRepository
}

func NewGenerateToken(
	config *config.Config,
	tokenMaker token.Maker,
	userRepository repository.UserRepository,
	sessionRepository repository.SessionRepository,
) *GenerateToken {
	return &GenerateToken{
		config:            config,
		tokenMaker:        tokenMaker,
		userRepository:    userRepository,
		sessionRepository: sessionRepository,
	}
}

func (g *GenerateToken) Execute(ctx context.Context, input *GenerateTokenInput) (output *GenerateTokenOutput, err error) {
	user, err := g.userRepository.GetUserByUsername(ctx, input.Username)
	if err != nil {
		fmt.Println("user not found", err.Error())
		return
	}

	err = domain.CheckPassword(input.Password, user.Password())
	if err != nil {
		fmt.Println("check user password error", err.Error())
		return
	}

	accessToken, err := g.tokenMaker.CreateToken(input.Username, g.config.AccessTokenDuration)
	if err != nil {
		fmt.Println("error generating access token", err.Error())
		return
	}

	refreshToken, err := g.tokenMaker.CreateToken(input.Username, g.config.RefreshTokenDuration)
	if err != nil {
		fmt.Println("error generating refresh token", err.Error())
		return
	}

	session, err := domain.NewSession(
		user.Username(),
		refreshToken,
		"userAgent",
		"clientIp",
		false,
		time.Now().Add(g.config.RefreshTokenDuration),
	)
	if err != nil {
		fmt.Println("error generating session", err.Error())
		return
	}

	err = g.sessionRepository.Save(ctx, session)
	if err != nil {
		fmt.Println("error saving session", err.Error())
		return
	}

	output = &GenerateTokenOutput{
		SessionID:             session.ID().String(),
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  time.Now().Add(g.config.AccessTokenDuration),
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: session.ExpiresAt(),
	}

	return
}
