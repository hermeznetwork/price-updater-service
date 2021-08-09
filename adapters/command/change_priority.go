package command

import (
	"github.com/hermeznetwork/price-updater-service/core/ports"
	"github.com/hermeznetwork/price-updater-service/core/services"
)

type ChangePriorityCommand struct {
	Priority string
	srv      *services.ConfigUpdaterService
}

func NewChangePriorityCommand(priority string, srv *services.ConfigUpdaterService) ports.Command {
	return &ChangePriorityCommand{
		Priority: priority,
		srv:      srv,
	}
}

func (cmd *ChangePriorityCommand) Execute() error {
	return cmd.srv.ChangePriority(cmd.Priority)
}
