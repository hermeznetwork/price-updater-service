package services

import (
	"context"

	"github.com/hermeznetwork/price-updater-service/core/domain"
	"github.com/hermeznetwork/price-updater-service/core/ports"
)

type FiatUpdaterService struct {
	fr ports.FiatPriceRepository
	pr ports.FiatProvider
}

func NewFiatUpdaterServices(fr ports.FiatPriceRepository, pr ports.FiatProvider) *FiatUpdaterService {
	return &FiatUpdaterService{
		fr: fr,
		pr: pr,
	}
}

func (f *FiatUpdaterService) GetPrices() ([]domain.FiatPrice, error) {
	ctx := context.Background()
	return f.fr.GetFiatPrices(ctx, "USD")
}

func (f *FiatUpdaterService) UpdatePrices() error {
	ctx := context.Background()

	// get prices from fiat API.
	rates, err := f.pr.GetPrices(ctx)
	if err != nil {
		return err
	}

	// get all price from database with baseCurrency DB
	currencies, err := f.fr.GetFiatPrices(ctx, "USD")

	for currency, price := range rates {
		// check if currency already exists
		var exist bool
		for i := 0; i < len(currencies); i++ {
			if currency == currencies[i].Currency {
				exist = true
				break
			}
		}

		if exist {
			// TODO: No log yet
			err = f.fr.UpdateFiatPrice(ctx, "USD", currency, price)
		} else {
			err = f.fr.CreateFiatPrice(ctx, "USD", currency, price)
		}
	}
	return err
}
