package command

import (
	"log"

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
	log.Println("Executing UpdatePrices()")
	return cmd.srv.LoadAndExecutePriceProvider()
}
