package cli

import (
	"github.com/hermeznetwork/hermez-node/log"
	"os"

	"github.com/hermeznetwork/price-updater-service/adapters/bbolt"
	"github.com/hermeznetwork/price-updater-service/config"
	"github.com/hermeznetwork/price-updater-service/core/services"
	"github.com/spf13/cobra"
)

var origins string
var apiKey string

var setupOriginCmd = &cobra.Command{
	Use:   "setup-origin",
	Short: "setup origins allowed to access API, comma separated",
	Run: func(cmd *cobra.Command, args []string) {
		setupOrigin(cfg)
	},
}

var setupAPIKeyCmd = &cobra.Command{
	Use:   "setup-apikey",
	Short: "setup apikey allowed to access API",
	Run: func(cmd *cobra.Command, args []string) {
		setupAPIKey(cfg)
	},
}

func init() {
	setupOriginCmd.Flags().StringVar(&origins, "origins", "*", "origin list comma separated")
	setupAPIKeyCmd.Flags().StringVar(&apiKey, "apiKey", "pr1c3upd4t3r", "api key used to access endpoints")
}

func setupOrigin(cfg config.Config) {
	log.Info("set up origin list")

	bboltCon := bbolt.NewConnection(cfg.Bbolt)
	projectConfigRepository := bbolt.NewProjectConfigRepository(bboltCon)
	projectConfigService := services.NewProjectConfigServices(projectConfigRepository)
	if err := projectConfigService.SaveOriginValue(origins); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}

func setupAPIKey(cfg config.Config) {
	log.Info("set up apikey")

	bboltCon := bbolt.NewConnection(cfg.Bbolt)
	projectConfigRepository := bbolt.NewProjectConfigRepository(bboltCon)
	projectConfigService := services.NewProjectConfigServices(projectConfigRepository)
	if err := projectConfigService.SaveAPIKey(apiKey); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}
