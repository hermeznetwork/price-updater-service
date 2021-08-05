package cli

import (
	"log"
	"os"

	"github.com/hermeznetwork/price-updater-service/adapters/bbolt"
	"github.com/hermeznetwork/price-updater-service/adapters/command"
	"github.com/hermeznetwork/price-updater-service/core/services"
	"github.com/spf13/cobra"
)

var priority string

var priorityCmd = &cobra.Command{
	Use:   "change-priority",
	Short: "command to change the provider priority",
	Run: func(cmd *cobra.Command, args []string) {
		changePriority(priority)
	},
}

func init() {
	priorityCmd.Flags().StringVar(&priority, "priority", "", "priority provider")

	priorityCmd.MarkFlagRequired("priority")
}

func changePriority(priority string) {
	log.Println("changing provider priority", priority)
	bboltConn := bbolt.NewConnection(cfg.Bbolt)
	priorityRepository := bbolt.NewProviderConfigRepository(bboltConn)
	priorityServices := services.NewConfigUpdaterServices(priorityRepository)
	changePriorityCmd := command.NewChangePriorityCommand(priority, priorityServices)
	err := changePriorityCmd.Execute()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
