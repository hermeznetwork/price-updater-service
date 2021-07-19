package ports

import (
	"context"

	"github.com/hermeznetwork/price-updater-service/core/domain"
)

type PriceProvider interface {
	GetPrices(ctx context.Context) ([]map[uint]float64, error)
}

type ConfigProviderRepository interface {
	CurrentProvider() (string, error)
	ChangeRunningProvider(provider string) error
	SaveConfig(provider string, data domain.PriceProvider) error
	LoadConfig(provider string) (domain.PriceProvider, error)
}
