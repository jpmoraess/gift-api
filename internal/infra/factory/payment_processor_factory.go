package factory

import (
	"github.com/jpmoraess/gift-api/internal/infra/chain"
	"github.com/jpmoraess/gift-api/internal/infra/gateway"
)

type PaymentProcessorFactory struct {
	asaasGateway *gateway.Asaas
}

func NewPaymentProcessorFactory(asaasGateway *gateway.Asaas) *PaymentProcessorFactory {
	return &PaymentProcessorFactory{asaasGateway: asaasGateway}
}

func (factory *PaymentProcessorFactory) CreatePaymentProcessor() chain.PaymentProcessor {
	asaasProcessor := chain.NewAsaasPaymentProcessor(factory.asaasGateway, nil)
	return asaasProcessor
}
