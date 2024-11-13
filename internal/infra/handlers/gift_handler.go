package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jpmoraess/gift-api/internal/application/usecase"
)

type GiftHandler struct {
	createGift *usecase.CreateGift
}

func NewGiftHandler(createGift *usecase.CreateGift) *GiftHandler {
	return &GiftHandler{createGift: createGift}
}

// CreateGift - handles the creation of a new gift
//
//	@Summary		Create a new gift
//	@Description	Create a new gift
//	@Tags			gifts
//	@Accept			json
//	@Produce		json
//	@Param			request	body	usecase.CreateGiftInput	true	"the request body for gift creation"
//	@Router			/v1/gifts [post]
func (handler *GiftHandler) CreateGift(c *fiber.Ctx) (err error) {
	input := new(usecase.CreateGiftInput)
	if err = c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	output, err := handler.createGift.Execute(c.Context(), input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(output)
}
