package command

import (
	"github.com/hermeznetwork/price-updater-service/core/ports"
	"github.com/hermeznetwork/price-updater-service/core/services"
)

type ChangeProviderCommand struct {
	Provider string
	srv      *services.ConfigUpdaterService
}

func NewChangeProviderCommand(name string, srv *services.ConfigUpdaterService) ports.Command {
	return &ChangeProviderCommand{
		Provider: name,
		srv:      srv,
	}
}

func (c *ChangeProviderCommand) Execute() error {
	return c.srv.ChangeRunningProvider(c.Provider)
}
