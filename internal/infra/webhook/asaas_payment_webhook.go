package webhook

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jpmoraess/gift-api/config"
)

type AsaasPaymentWebhook struct {
	config *config.Config
}

func NewAsaasPaymentWebhook(config *config.Config) *AsaasPaymentWebhook {
	return &AsaasPaymentWebhook{config: config}
}

func (wh *AsaasPaymentWebhook) HandlePaymentEvent(c *fiber.Ctx) (err error) {
	// 1. receber token
	token := c.Get("asaas-access-token")
	_ = token

	// 2. validar token

	// 3. parsear corpo da requisição
	event := new(AsaasPaymentEvent)
	if err = c.BodyParser(event); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// 4. processar evento idempotência idempotência idempotência idempotência

	return c.SendStatus(fiber.StatusNoContent)
}
