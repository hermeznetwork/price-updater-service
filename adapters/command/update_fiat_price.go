package command

import (
	"log"

	"github.com/hermeznetwork/price-updater-service/core/services"
)

type UpdateFiatPriceCommand struct {
	srv *services.FiatUpdaterService
}

func NewUpdateFiatPriceCommand(srv *services.FiatUpdaterService) *UpdateFiatPriceCommand {
	return &UpdateFiatPriceCommand{
		srv: srv,
	}
}

func (cmd *UpdateFiatPriceCommand) Execute() error {
	log.Println("Executing Fiat UpdatePrices()")
	return cmd.srv.UpdatePrices()
}
