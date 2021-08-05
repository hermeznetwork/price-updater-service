package services

import "log"

type ProviderUpdateOrchestratorService struct {
	providerSelector *ProviderSelectorService
	priceUpdater     *PriceUpdaterService
	fiatUpdater      *FiatUpdaterService
}

func NewPriceUpdateOrchestratorService(ps *ProviderSelectorService, pus *PriceUpdaterService, fus *FiatUpdaterService) *ProviderUpdateOrchestratorService {
	return &ProviderUpdateOrchestratorService{
		providerSelector: ps,
		priceUpdater:     pus,
		fiatUpdater:      fus,
	}
}

func (orchestrator *ProviderUpdateOrchestratorService) UpdatePrices() error {
	log.Println("Executing provider update")
	err := orchestrator.priceUpdater.UpdatePrices()
	if err != nil {
		return err
	}

	log.Println("Executing fiat update")
	return orchestrator.fiatUpdater.UpdatePrices()
}

func (orchestrator *ProviderUpdateOrchestratorService) LoadAndExecutePriceProvider() error {
	log.Println("Check if provider has changed")
	providers, err := orchestrator.providerSelector.AllProviders()
	if err != nil {
		log.Println("fail to retrive current provider: ", err.Error())
		return err
	}
	orchestrator.priceUpdater.LoadProvider(providers)
	return orchestrator.UpdatePrices()
}
