package services

import "github.com/hermeznetwork/price-updater-service/core/ports"

type PriceUpdaterService struct {
}

func NewPriceUpdaterService(provider ports.PriceProvider) *PriceUpdaterService {
	return &PriceUpdaterService{}
}

func (s *PriceUpdaterService) UpdatePrices() {
	panic("not implemented yet")
}

