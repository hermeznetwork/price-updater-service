package services

import (
	"github.com/hermeznetwork/price-updater-service/core/domain"
	"github.com/hermeznetwork/price-updater-service/core/ports"
)

type ConfigUpdaterService struct {
	cpr ports.ConfigProviderRepository
}

func NewConfigUpdaterServices(cpr ports.ConfigProviderRepository) *ConfigUpdaterService {
	return &ConfigUpdaterService{
		cpr: cpr,
	}
}

func (c *ConfigUpdaterService) LoadConfig(provider string) (domain.PriceProvider, error) {
	return c.cpr.LoadConfig(provider)
}

func (c *ConfigUpdaterService) UpdateConfig(pr *domain.PriceProvider) error {
	return c.cpr.SaveConfig(pr.Provider, *pr)
}
