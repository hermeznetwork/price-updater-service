package ports

import (
	"context"

	"github.com/hermeznetwork/price-updater-service/core/domain"
)

type FiatPriceRepository interface {
	GetFiatPrices(ctx context.Context) ([]domain.FiatPrice, error)
	CreateFiatPrice(ctx context.Context, fp domain.FiatPrice) error
}
