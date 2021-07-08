package ports

import "context"

// import "context"

type PriceProvider interface {
	GetPrices(ctx context.Context) (float64, error)
}
