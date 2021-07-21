package services

type UpdateOrchestratorService struct {
	pus *PriceUpdaterService
	fus *FiatUpdaterService
}

func NewOrchestratorService(pus *PriceUpdaterService, fus *FiatUpdaterService) *UpdateOrchestratorService {
	return &UpdateOrchestratorService{
		pus: pus,
		fus: fus,
	}
}

func (orchestrator *UpdateOrchestratorService) UpdatePrices() error {
	err := orchestrator.pus.UpdatePrices()
	if err != nil {
		return err
	}

	return orchestrator.fus.UpdatePrices()
}
