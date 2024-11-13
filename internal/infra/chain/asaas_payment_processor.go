package chain

import (
	"context"
	"fmt"
	"github.com/jpmoraess/gift-api/internal/infra/gateway"
)

type AsaasPaymentProcessor struct {
	BasePaymentProcessor
	gateway *gateway.Asaas
}

func NewAsaasPaymentProcessor(gateway *gateway.Asaas, next PaymentProcessor) *AsaasPaymentProcessor {
	processor := &AsaasPaymentProcessor{
		gateway:              gateway,
		BasePaymentProcessor: BasePaymentProcessor{next: next},
	}
	return processor
}

func (a *AsaasPaymentProcessor) ProcessPayment(ctx context.Context, input *ProcessPaymentInput) (output *ProcessPaymentOutput, err error) {
	fmt.Printf("processing payment through Asaas, %+v", input)

	request := &gateway.CreateBillingRequest{
		Customer:    "6348759",
		BillingType: gateway.Pix,
		Value:       input.Amount,
		DueDate:     "2024-11-14",
	}

	response, err := a.gateway.CreateBilling(ctx, request)
	if err != nil {
		fmt.Println("error creating billing request:", err)
		return nil, err
	}
	fmt.Printf("payment successfully: %+v\n", response)

	if len(response.ID) > 0 {
		fmt.Println("payment id:", response.ID)
		output = &ProcessPaymentOutput{
			ID: response.ID,
		}
		return output, err
	} else {
		fmt.Println("failed processing payment request through Asaas...")
		if a.next == nil {
			return nil, fmt.Errorf("no next payment processor has been provided")
		}
		return a.next.ProcessPayment(ctx, input)
	}
}
