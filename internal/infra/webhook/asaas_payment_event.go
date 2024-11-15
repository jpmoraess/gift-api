package webhook

// See: https://docs.asaas.com/docs/webhook-para-cobrancas

type PaymentEventType string

const (
	// Geração de nova cobrança
	PAYMENT_CREATED PaymentEventType = "PAYMENT_CREATED"

	// Cobrança confirmada (pagamento efetuado, porém o saldo ainda não foi disponibilizado)
	PAYMENT_CONFIRMED PaymentEventType = "PAYMENT_CONFIRMED"

	// Cobrança recebida
	PAYMENT_RECEIVED PaymentEventType = "PAYMENT_RECEIVED"
)

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
