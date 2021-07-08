package services

import "github.com/hermeznetwork/price-updater-service/core/ports"

type ProviderSelectorService struct {
}

func NewProviderSelectorService(providers []ports.PriceProvider) *ProviderSelectorService {
	return &ProviderSelectorService{}
}

func (s *ProviderSelectorService) Select(name string) ports.PriceProvider {
	panic("not implemented yet")
}
