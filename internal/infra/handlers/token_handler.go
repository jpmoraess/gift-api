package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jpmoraess/gift-api/internal/application/usecase"
)

type TokenHandler struct {
	generateToken *usecase.GenerateToken
}

func NewTokenHandler(generateToken *usecase.GenerateToken) *TokenHandler {
	return &TokenHandler{generateToken: generateToken}
}

// GenerateToken - handles the creation of a new access token
//
//	@Summary		Generate a new access token
//	@Description	Generate a new access token
//	@Tags			token
//	@Accept			json
//	@Produce		json
//	@Param			request	body	usecase.GenerateTokenInput	true	"the request body for token generation"
//	@Router			/auth/token [post]
func (handler *TokenHandler) GenerateToken(c *fiber.Ctx) (err error) {
	input := new(usecase.GenerateTokenInput)
	if err = c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	output, err := handler.generateToken.Execute(c.Context(), input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(output)
}
