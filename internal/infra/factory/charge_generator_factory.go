package factory

import (
	"github.com/jpmoraess/gift-api/internal/infra/chain"
	"github.com/jpmoraess/gift-api/internal/infra/gateway"
)

type ChargeGeneratorFactory struct {
	asaasGateway *gateway.AsaasGateway
}

func NewChargeGeneratorFactory(asaasGateway *gateway.AsaasGateway) *ChargeGeneratorFactory {
	return &ChargeGeneratorFactory{asaasGateway: asaasGateway}
}

func (factory *ChargeGeneratorFactory) CreateChargeGeneratorChain() chain.ChargeGenerator {
	asaas := chain.NewAsaasChargeGenerator(factory.asaasGateway, nil)
	return asaas
}
