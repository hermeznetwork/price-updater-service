package cli

import (
	"github.com/hermeznetwork/hermez-node/log"
	"os"

	"github.com/hermeznetwork/price-updater-service/adapters/bbolt"
	"github.com/hermeznetwork/price-updater-service/adapters/command"
	"github.com/hermeznetwork/price-updater-service/config"
	"github.com/hermeznetwork/price-updater-service/core/services"
	"github.com/spf13/cobra"
)

var changeProviderName string

var changeProviderCmd = &cobra.Command{
	Use:   "change-provider",
	Short: "command to change a running provider",
	Run: func(cmd *cobra.Command, args []string) {
		changeProvider(cfg)
	},
}

func init() {
	changeProviderCmd.Flags().StringVar(&changeProviderName, "provider", "", "provider name")

	changeProviderCmd.MarkFlagRequired("provider")
}

func changeProvider(cfg config.Config) {
	log.Info("running change provider")
	bboltCon := bbolt.NewConnection(cfg.Bbolt)
	configProviderRepository := bbolt.NewProviderConfigRepository(bboltCon)
	configUpdaterServices := services.NewConfigUpdaterServices(configProviderRepository)
	changeProviderCmd := command.NewChangeProviderCommand(changeProviderName, configUpdaterServices)
	err := changeProviderCmd.Execute()
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	log.Info("provider changed!")
}
