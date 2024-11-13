package asaas

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Asaas struct {
	baseUrl string
	http    *http.Client
}

type BillingType string

const (
	Pix        BillingType = "PIX"
	Boleto     BillingType = "BOLETO"
	CreditCard BillingType = "CREDIT_CARD"
)

type CreateBillingRequest struct {
	Customer    string      `json:"customer"`
	BillingType BillingType `json:"billingType"`
	Value       float64     `json:"value"`
	DueDate     string      `json:"dueDate"`
}

type CreateBillingResponse struct {
	ID          string  `json:"id"`
	CreatedDate string  `json:"dateCreated"`
	Value       float64 `json:"value"`
}

func NewAsaas(baseUrl string, http *http.Client) *Asaas {
	return &Asaas{
		baseUrl: baseUrl,
		http:    http,
	}
}

// CreateBilling - criação de cobrança
func (a *Asaas) CreateBilling(
	ctx context.Context,
	token string,
	request *CreateBillingRequest,
) (response *CreateBillingResponse, err error) {

	url := fmt.Sprintf("%s/v3/lean/payments", a.baseUrl)

	body, err := json.Marshal(request)
	if err != nil {
		return
	}

	b := bytes.NewReader(body)

	// criando a requisição
	req, err := http.NewRequestWithContext(ctx, "POST", url, b)
	if err != nil {
		return
	}
	req.Header.Add("access_token", token)
	req.Header.Add("User-Agent", "i-gift")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	// enviando a requisição
	resp, err := a.http.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// lendo o corpo da resposta
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	fmt.Println("status code: ", resp.StatusCode)

	// deserializando response body para a sruct CreateBillingResponse
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		fmt.Println("erro ao deserializar a resposta: ", err)
		return
	}

	return
}
