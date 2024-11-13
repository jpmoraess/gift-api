package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jpmoraess/gift-api/internal/application/usecase"
)

type TransactionHandler struct {
	processPayment *usecase.ProcessPayment
}

func NewTransactionHandler(processPayment *usecase.ProcessPayment) *TransactionHandler {
	return &TransactionHandler{
		processPayment: processPayment,
	}
}

// ProcessPayment - handles the payment processor
//
//	@Summary		Create a new payment
//	@Description	Create a new payment
//	@Tags			gifts
//	@Accept			json
//	@Produce		json
//	@Param			request	body	usecase.ProcessPaymentInput	true	"the request body for transaction creation"
//	@Router			/v1/transactions [post]
func (handler *TransactionHandler) ProcessPayment(c *fiber.Ctx) error {
	input := new(usecase.ProcessPaymentInput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	output, err := handler.processPayment.Execute(c.Context(), input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(output)
}
