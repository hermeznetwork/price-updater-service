package ports

import "context"

type PriceProvider interface {
	GetPrices(ctx context.Context) ([]map[uint]float64, error)
}
