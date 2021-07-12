package services

import (
	"context"
	"log"

	"github.com/hermeznetwork/price-updater-service/core/domain"
	"github.com/hermeznetwork/price-updater-service/core/ports"
)

type PriceUpdaterService struct {
	pr  ports.PriceProvider
	tr  ports.TokenRepository
	ctx context.Context
}

func NewPriceUpdaterService(provider ports.PriceProvider, tr ports.TokenRepository, ctx context.Context) *PriceUpdaterService {
	return &PriceUpdaterService{
		pr:  provider,
		tr:  tr,
		ctx: ctx,
	}
}

func (s *PriceUpdaterService) GetTokens() ([]domain.Token, error) {
	return s.tr.GetTokens(s.ctx)
}

func (s *PriceUpdaterService) UpdatePrices() error {
	// get prices from provider
	prices, err := s.pr.GetPrices(s.ctx)
	if err != nil {
		return err
	}

	for _, price := range prices {
		for tokenID, value := range price {
			err = s.tr.UpdateTokenPrice(s.ctx, tokenID, value)
			if err != nil {
				log.Println("try update token price: ", err.Error())
			}
		}
	}

	return nil
}
