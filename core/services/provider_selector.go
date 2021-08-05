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

func (selector *ProviderSelectorService) CurrentProvider() (ports.PriceProvider, error) {
	provider, err := selector.brepo.CurrentProvider()
	if err != nil {
		log.Println("try get the current provider: ", err.Error())
		return nil, err
	}
	log.Println("Get the current provider:", provider)

	return selector.Select(provider)
}

func (selector *ProviderSelectorService) Select(name string) (ports.PriceProvider, error) {
	configProvidiver, err := selector.brepo.LoadConfig(name)
	if err != nil {
		return nil, err
	}
	switch configProvidiver.Provider {
	case "bitfinex":
		return bitfinex.NewClient(configProvidiver.MappingBetweenNetwork), nil
	case "coingecko":
		return coingecko.NewClient(configProvidiver.MappingBetweenNetwork), nil
	case "uniswap":
		return uniswap.NewClient(selector.cfg.EthConfig)
	default:
		log.Println("provider not found")
		return nil, nil
	}
}
func (selector *ProviderSelectorService) AllProviders() ([]ports.PriceProvider, error) {
	var err error
	providers := []string{"bitfinex","coingecko","uniswap"} //The order is prioritized
	log.Println("Get all providers:", providers)
	var priceProviders []ports.PriceProvider
	for i:=0;i<len(providers);i++{
		var prov ports.PriceProvider
		prov, err = selector.Select(providers[i])
		if err != nil {
			log.Println("Error getting provider "+providers[i]+": ",err)
		}
		priceProviders = append(priceProviders, prov)
	}
	return priceProviders, err
}