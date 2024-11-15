package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jpmoraess/gift-api/internal/application/usecase"
)

type UserHandler struct {
	createUser *usecase.CreateUser
}

func NewUserHandler(createUser *usecase.CreateUser) *UserHandler {
	return &UserHandler{createUser: createUser}
}

// CreateUser - handles the creation of a new user
//
//	@Summary		Create a new user
//	@Description	Create a new user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			request	body	usecase.CreateUserInput	true	"the request body for user creation"
//	@Router			/v1/users [post]
func (handler *UserHandler) CreateUser(c *fiber.Ctx) error {
	input := new(usecase.CreateUserInput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	output, err := handler.createUser.Execute(c.Context(), input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(output)
}
