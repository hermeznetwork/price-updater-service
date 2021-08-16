package cli

import (
	"os"

	"github.com/hermeznetwork/hermez-node/log"

	"github.com/hermeznetwork/price-updater-service/adapters/bbolt"
	"github.com/hermeznetwork/price-updater-service/config"
	"github.com/hermeznetwork/price-updater-service/core/services"
	"github.com/spf13/cobra"
)

var origins string

var setupOriginCmd = &cobra.Command{
	Use:   "setup-origin",
	Short: "setup origins allowed to access API, comma separated",
	Run: func(cmd *cobra.Command, args []string) {
		setupOrigin(cfg)
	},
}

func init() {
	setupOriginCmd.Flags().StringVar(&origins, "origins", "*", "origin list comma separated")
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
	log.Info("origin list updated!")
}
