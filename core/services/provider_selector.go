package services

import (
	"context"

	"github.com/hermeznetwork/hermez-node/log"
	"github.com/hermeznetwork/price-updater-service/core/domain"

	"github.com/hermeznetwork/price-updater-service/adapters/bitfinex"
	"github.com/hermeznetwork/price-updater-service/adapters/coingecko"
	"github.com/hermeznetwork/price-updater-service/adapters/uniswap"
	"github.com/hermeznetwork/price-updater-service/config"
	"github.com/hermeznetwork/price-updater-service/core/ports"
)

type ProviderSelectorService struct {
	brepo ports.ConfigProviderRepository
	cfg   config.Config
	tr    ports.TokenRepository
}

func NewProviderSelectorService(brepo ports.ConfigProviderRepository, cfg config.Config, priceRepository ports.TokenRepository) *ProviderSelectorService {
	return &ProviderSelectorService{
		brepo: brepo,
		cfg:   cfg,
		tr: priceRepository,
	}
}

func (selector *ProviderSelectorService) CurrentProvider() (ports.PriceProvider, error) {
	provider, err := selector.brepo.CurrentProvider()
	if err != nil {
		log.Error("try get the current provider: ", err.Error())
		return nil, err
	}
	log.Info("Get the current provider:", provider)

	return selector.Select(provider, []domain.Token{})
}

func (selector *ProviderSelectorService) Select(name string, defaultTokensInfo []domain.Token) (ports.PriceProvider, error) {
	configProvidiver, err := selector.brepo.LoadConfig(name)
	if err != nil {
		return nil, err
	}
	switch configProvidiver.Provider {
	case "bitfinex":
		allTokens := mergeTokens(defaultTokensInfo, configProvidiver.MappingBetweenNetwork, name)
		return bitfinex.NewClient(allTokens), nil
	case "coingecko":
		allTokens := mergeTokens(defaultTokensInfo, configProvidiver.MappingBetweenNetwork, name)
		return coingecko.NewClient(allTokens), nil
	case "uniswap":
		return uniswap.NewClient(selector.cfg.EthConfig)
	default:
		log.Error("provider not found")
		return nil, nil
	}
}
func (selector *ProviderSelectorService) PrioritizedProviders(ctx context.Context) ([]ports.PriceProvider, error) {
	var err error
	//Get priority from db
	priorityProviders, err := selector.brepo.PriorityProviders()
	if err != nil {
		log.Error("try get the current provider: ", err.Error())
		return nil, err
	}
	//Get default token info
	defaultTokensInfo, err := selector.tr.GetTokens(ctx, 0, 0, "ASC")
	if err != nil {
		log.Error("error getting default token values from db: ", err)
	}
	var priceProviders []ports.PriceProvider
	for i := 0; i < len(priorityProviders); i++ {
		var prov ports.PriceProvider
		prov, err = selector.Select(priorityProviders[i], defaultTokensInfo)
		if err != nil {
			log.Error("Error getting provider "+priorityProviders[i]+": ", err)
		}
		priceProviders = append(priceProviders, prov)
	}
	return priceProviders, err
}

func mergeTokens(defaultData []domain.Token, confData map[uint]string, provider string) map[uint]string {
	for i := 0; i < len(defaultData); i++ {
		var flag bool
        for k, _ := range confData {
			if k == defaultData[i].ID {
				flag = true
			}
		}
		if !flag {
			//Add token info to the mapping
			if provider == "bitfinex" {
				confData[defaultData[i].ID] = defaultData[i].Symbol
			} else if provider == "coingecko" {
				confData[defaultData[i].ID] = defaultData[i].Address
			}
		}
	}
	return confData
}
