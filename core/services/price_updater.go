package services

import (
	"context"
	"log"

	"github.com/hermeznetwork/price-updater-service/core/domain"
	"github.com/hermeznetwork/price-updater-service/core/ports"
)

type PriceUpdaterService struct {
	pr  []ports.PriceProvider
	tr  ports.TokenRepository
	ctx context.Context
}

func NewPriceUpdaterService(ctx context.Context, providers []ports.PriceProvider, tr ports.TokenRepository) *PriceUpdaterService {
	return &PriceUpdaterService{
		pr:  providers,
		tr:  tr,
		ctx: ctx,
	}
}

func (s *PriceUpdaterService) LoadProvider(providers []ports.PriceProvider) {
	s.pr = providers
}

func (s *PriceUpdaterService) GetToken(tokenID uint) (domain.Token, error) {
	return s.tr.GetToken(s.ctx, tokenID)
}

func (s *PriceUpdaterService) GetTokens(fromItem uint, limit uint, order string) ([]domain.Token, error) {
	return s.tr.GetTokens(s.ctx, fromItem, limit, order)
}

func (s *PriceUpdaterService) UpdatePrices() error {
	// get prices from providers
	var (
		prices []map[uint]float64
		tokenErrs []uint
		err error
	)
	for i:=0; i<len(s.pr); i++{
		if len(tokenErrs) == 0 {
			prices, tokenErrs, err = s.pr[i].GetPrices(s.ctx)
			log.Println("tokenErrs: ",tokenErrs)
			if err != nil || len(tokenErrs) != 0 {
				log.Println("Error getting prices from provider: ", err, " TokenErrs length is ", len(tokenErrs))
				continue
			}
		} else {
			prices, tokenErrs, err = s.pr[i].GetFailedPrices(s.ctx, prices, tokenErrs)
			log.Println("tokenErrs: ",tokenErrs)
			if err != nil || len(tokenErrs) != 0 {
				log.Println("Error getting prices from provider: ", err, " TokenErrs length is ", len(tokenErrs))
				continue
			}
		}
		log.Println("Prices: ",prices)
		break
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

func (s *PriceUpdaterService) UpdatePrice(tokenID uint, value float64) error {
	return s.tr.UpdateTokenPrice(s.ctx, tokenID, value)
}
