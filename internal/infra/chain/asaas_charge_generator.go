package chain

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jpmoraess/gift-api/internal/infra/gateway"
)

type AsaasChargeGenerator struct {
	BaseChargeGenerator
	gateway *gateway.AsaasGateway
}

func NewAsaasChargeGenerator(gateway *gateway.AsaasGateway, next ChargeGenerator) *AsaasChargeGenerator {
	processor := &AsaasChargeGenerator{
		gateway:             gateway,
		BaseChargeGenerator: BaseChargeGenerator{next: next},
	}
	return processor
}

func (a *AsaasChargeGenerator) GenerateCharge(ctx context.Context, input *GenerateChargeInput) (output *GenerateChargeOutput, err error) {
	log.Printf("generating charge through AsaasGateway for input: %+v\n", input)

	request := &gateway.CreatePaymentRequest{
		Customer:    "6348759",
		BillingType: gateway.Pix,
		Value:       input.Amount,
		DueDate:     time.Now().Format("2006-01-02"), // TODO: due date, valid date, week
	}

	response, err := a.gateway.CreatePayment(ctx, request)
	if err != nil {
		log.Printf("error creating payment: %+v\n", err)
		return nil, err
	}

	if len(response.ID) > 0 {
		output = &GenerateChargeOutput{ID: response.ID}
		return output, err
	} else {
		if a.next == nil {
			return nil, fmt.Errorf("no next charge generator has been provided")
		}
		return a.next.GenerateCharge(ctx, input)
	}
}
