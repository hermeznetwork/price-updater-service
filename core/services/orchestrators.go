package services

import "github.com/hermeznetwork/hermez-node/log"

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
	log.Info("Executing provider update")
	err := orchestrator.priceUpdater.UpdatePrices()
	if err != nil {
		return err
	}

	log.Info("Executing fiat update")
	return orchestrator.fiatUpdater.UpdatePrices()
}

func (orchestrator *ProviderUpdateOrchestratorService) LoadAndExecutePriceProvider() error {
	log.Info("Check if provider has changed")
	providers, err := orchestrator.providerSelector.PrioritizedProviders(orchestrator.priceUpdater.ctx, orchestrator.priceUpdater.tr)
	if err != nil {
		log.Error("fail to retrive current provider: ", err.Error())
		return err
	}
	orchestrator.priceUpdater.LoadProviders(providers)
	return orchestrator.UpdatePrices()
}
