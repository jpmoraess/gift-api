package gateway

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jpmoraess/gift-api/config"
)

type AsaasGateway struct {
	config *config.Config
	http   *http.Client
}

// BillingType
type BillingType string

const (
	Pix        BillingType = "PIX"
	Boleto     BillingType = "BOLETO"
	CreditCard BillingType = "CREDIT_CARD"
)

type CreatePaymentCallback struct {
	// SuccessURL - represents the URL that the customer will be redirected to after successful chain of the invoice or chain link
	SuccessURL string `json:"successUrl"`
	// AutoRedirect - Define whether the customer will be automatically redirected or will just be informed with a button to return to the website. The default is true, if you want to disable it, enter false
	AutoRedirect bool `json:"autoRedirect"`
}

// CreatePaymentRequest - Object to create billing
// See: https://docs.asaas.com/reference/criar-nova-cobranca-com-dados-resumidos-na-resposta
type CreatePaymentRequest struct {
	// Customer - represents the unique customer identifier in Asaas
	Customer string `json:"customer" validate:"required"`
	// BillingType - represents the chain method (PIX, BOLETO, CREDIT_CARD)
	BillingType BillingType `json:"billingType" validate:"required"`
	// Value - represents the charge amount, 0.0
	Value float64 `json:"value" validate:"required"`
	// DueDate - represents the due date, format: yyyy-MM-dd
	DueDate string `json:"dueDate" validate:"required"`
	// Description - represents the charge description
	Description string `json:"description"`
}

type CreatePaymentResponse struct {
	ID          string  `json:"id"`
	CreatedDate string  `json:"dateCreated"`
	Value       float64 `json:"value"`
}

type ErrorResponse struct {
	Errors []struct {
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"errors"`
}

func NewAsaasGateway(config *config.Config, http *http.Client) *AsaasGateway {
	return &AsaasGateway{
		config: config,
		http:   http,
	}
}

// CreatePayment - create payment
// See: https://docs.asaas.com/reference/criar-nova-cobranca-com-dados-resumidos-na-resposta
func (a *AsaasGateway) CreatePayment(ctx context.Context, request *CreatePaymentRequest) (response *CreatePaymentResponse, err error) {
	url := fmt.Sprintf("%s/v3/lean/payments", a.config.AsaasUrl)

	body, err := json.Marshal(request)
	if err != nil {
		fmt.Println("error while marshalling request body")
		return
	}

	b := bytes.NewReader(body)

	// creating the payment request
	req, err := http.NewRequestWithContext(ctx, "POST", url, b)
	if err != nil {
		fmt.Println("error while creating a payment request:", err)
		return
	}
	req.Header.Add("access_token", a.config.AsaasApiKey)
	req.Header.Add("User-Agent", "i-gift")
	req.Header.Add("Content-Type", "application/json")

	// sending the request
	resp, err := a.http.Do(req)
	if err != nil {
		fmt.Println("error while executing a request:", err)
		return
	}
	defer resp.Body.Close()

	// reading the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error while reading the response body:", err)
		return
	}

	// checking if the response status code is an error (4xx, 5xx)
	if resp.StatusCode >= 400 {
		var errorResponse ErrorResponse
		err = json.Unmarshal(respBody, &errorResponse)
		if err != nil {
			fmt.Println("error while deserialize error response: ", err)
			return
		}
		return nil, fmt.Errorf("error while creating the payment: %+v", errorResponse.Errors)
	}

	// deserializing response body to CreateBillingResponse struct
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		fmt.Println("error while deserialize response:", err)
		return
	}

	return
}
