package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jpmoraess/gift-api/internal/application/usecase"
)

type TransactionHandler struct {
	generateCharge *usecase.GenerateCharge
}

func NewTransactionHandler(generateCharge *usecase.GenerateCharge) *TransactionHandler {
	return &TransactionHandler{
		generateCharge: generateCharge,
	}
}

// CreateTransaction - handles the transaction creation
//
//	@Summary		Create a new transaction
//	@Description	Create a new transaction
//	@Tags			transactions
//	@Accept			json
//	@Produce		json
//	@Param			request	body	usecase.GenerateChargeInput	true	"the request body for transaction creation"
//	@Router			/v1/transactions [post]
func (handler *TransactionHandler) CreateTransaction(c *fiber.Ctx) error {
	var input usecase.GenerateChargeInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	output, err := handler.generateCharge.Execute(c.Context(), &input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(output)
}
