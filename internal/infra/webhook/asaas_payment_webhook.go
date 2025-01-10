package webhook

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jpmoraess/gift-api/config"
)

// See: https://docs.asaas.com/docs/webhook-para-cobrancas
type AsaasPaymentEvent struct {
	ID          string `json:"id"`
	Event       string `json:"event"`
	DateCreated string `json:"dateCreated"`
	Payment     struct {
		Object      string  `json:"payment"`
		ID          string  `json:"id"`
		DateCreated string  `json:"dateCreated"`
		Customer    string  `json:"customer"`
		PaymentLink string  `json:"paymentLink"`
		Value       float64 `json:"value"`
		NetValue    float64 `json:"netValue"`
		BillingType string  `json:"billingType"`
		Status      string  `json:"status"`
	} `json:"payment"`
}

type PaymentEventType string

const (
	// Geração de nova cobrança
	PAYMENT_CREATED PaymentEventType = "PAYMENT_CREATED"

	// Cobrança confirmada (pagamento efetuado, porém o saldo ainda não foi disponibilizado)
	PAYMENT_CONFIRMED PaymentEventType = "PAYMENT_CONFIRMED"

	// Cobrança recebida
	PAYMENT_RECEIVED PaymentEventType = "PAYMENT_RECEIVED"
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
