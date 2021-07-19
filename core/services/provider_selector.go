package services

import (
	"log"

	"github.com/hermeznetwork/price-updater-service/adapters/bitfinex"
	"github.com/hermeznetwork/price-updater-service/adapters/coingecko"
	"github.com/hermeznetwork/price-updater-service/adapters/uniswap"
	"github.com/hermeznetwork/price-updater-service/config"
	"github.com/hermeznetwork/price-updater-service/core/ports"
)

type ProviderSelectorService struct {
	brepo ports.ConfigProviderRepository
	cfg   config.Config
}

func NewProviderSelectorService(brepo ports.ConfigProviderRepository, cfg config.Config) *ProviderSelectorService {
	return &ProviderSelectorService{
		brepo: brepo,
		cfg:   cfg,
	}
}

func (s *ProviderSelectorService) CurrentProvider() (ports.PriceProvider, error) {
	provider, err := s.brepo.CurrentProvider()
	if err != nil {
		log.Println("try get the current provider: ", err.Error())
		return nil, err
	}
	log.Println("Get the current provider:", provider)

	return s.Select(provider)
}

func (s *ProviderSelectorService) Select(name string) (ports.PriceProvider, error) {
	configProvidiver, err := s.brepo.LoadConfig(name)
	if err != nil {
		return nil, err
	}
	switch configProvidiver.Provider {
	case "bitfinex":
		return bitfinex.NewClient(configProvidiver.MappingBetweenNetwork), nil
	case "coingecko":
		return coingecko.NewClient(configProvidiver.MappingBetweenNetwork), nil
	case "uniswap":
		return uniswap.NewClient(s.cfg.EthConfig)
	default:
		log.Println("provider not found")
		return nil, nil
	}
}
