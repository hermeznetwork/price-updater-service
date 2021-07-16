package ports

import (
	"context"

	"github.com/hermeznetwork/price-updater-service/core/domain"
)

type FiatProvider interface {
	GetPrices(ctx context.Context) (map[string]float64, error)
}

type FiatPriceRepository interface {
	GetFiatPrice(ctx context.Context, currency string) (domain.FiatPrice, error)
	GetFiatPrices(ctx context.Context, baseCurrency string) ([]domain.FiatPrice, error)
	CreateFiatPrice(ctx context.Context, base_currency, currency string, price float64) error
	UpdateFiatPrice(ctx context.Context, base_currency, currency string, price float64) error
}
