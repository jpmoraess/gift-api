package factory

import (
	"github.com/jpmoraess/gift-api/internal/infra/chain"
	"github.com/jpmoraess/gift-api/internal/infra/gateway"
)

type PaymentProcessorFactory struct {
	asaasPaymentGateway *gateway.AsaasPaymentGateway
}

func NewPaymentProcessorFactory(asaasPaymentGateway *gateway.AsaasPaymentGateway) *PaymentProcessorFactory {
	return &PaymentProcessorFactory{asaasPaymentGateway: asaasPaymentGateway}
}

func (factory *PaymentProcessorFactory) CreatePaymentProcessor() chain.PaymentProcessor {
	asaasPaymentProcessor := chain.NewAsaasPaymentProcessor(factory.asaasPaymentGateway, nil)
	return asaasPaymentProcessor
}
