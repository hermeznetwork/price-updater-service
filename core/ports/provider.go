package ports

import (
	"context"

	"github.com/hermeznetwork/price-updater-service/core/domain"
)

type PriceProvider interface {
	GetPrices(ctx context.Context) ([]map[uint]float64, []uint, error)
	GetFailedPrices(ctx context.Context, prices []map[uint]float64, tokenErrs []uint) ([]map[uint]float64, []uint, error)
}

type ConfigProviderRepository interface {
	SaveConfig(provider string, data domain.PriceProvider) error
	LoadConfig(provider string) (domain.PriceProvider, error)
	ChangePriority(priority string) error
	PriorityProviders() ([]string, error)
}
