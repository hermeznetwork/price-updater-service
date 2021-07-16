package ports

import (
	"context"

	"github.com/hermeznetwork/price-updater-service/core/domain"
)

type PriceProvider interface {
	GetPrices(ctx context.Context) ([]map[uint]float64, error)
}

type ConfigProviderRepository interface {
	SaveConfig() error
	LoadConfig() (domain.PriceProvider, error)
}
