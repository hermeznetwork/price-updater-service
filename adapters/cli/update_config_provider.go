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

var jsonConfigFile string
var provider string

var updaterCmd = &cobra.Command{
	Use:   "update-config",
	Short: "command to update a provider configuration",
	Run: func(cmd *cobra.Command, args []string) {
		updateConfig(cfg)
	},
}

func init() {
	updaterCmd.Flags().StringVar(&provider, "provider", "", "provider name")
	updaterCmd.Flags().StringVar(&jsonConfigFile, "configFile", "", "path to provider config file")

	updaterCmd.MarkFlagRequired("provider")
	updaterCmd.MarkFlagRequired("configFile")
}

func updateConfig(cfg config.Config) {
	log.Info("starting update configuration to ", provider)
	bboltConn := bbolt.NewConnection(cfg.Bbolt)
	defer bboltConn.End()
	configProviderRepository := bbolt.NewProviderConfigRepository(bboltConn)
	configUpdaterServices := services.NewConfigUpdaterServices(configProviderRepository)
	updatConfigCmd := command.NewUpdateProviderConfigCommand(provider, jsonConfigFile, configUpdaterServices)
	err := updatConfigCmd.Execute()
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	log.Info("configuration updated!")
}
