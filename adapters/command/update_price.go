package command

import (
	"github.com/hermeznetwork/hermez-node/log"

	"github.com/hermeznetwork/price-updater-service/core/services"
)

type UpdatePriceCommand struct {
	srv *services.ProviderUpdateOrchestratorService
}

func NewUpdatePriceCommand(srv *services.ProviderUpdateOrchestratorService) *UpdatePriceCommand {
	return &UpdatePriceCommand{
		srv: srv,
	}
}

func (cmd *UpdatePriceCommand) Execute() error {
	log.Info("Executing UpdatePrices()")
	return cmd.srv.LoadAndExecutePriceProvider()
}
