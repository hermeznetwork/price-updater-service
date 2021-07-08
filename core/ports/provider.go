package ports

import "context"

// import "context"

type PriceProvider interface {
	GetPrices(ctx context.Context) (float64, error)
}

type PriceProviderRepository interface {
	//SavePrice(ctx context.Context) error
	//GetPrice(ctx context.Context) error
}
