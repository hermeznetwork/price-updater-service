package cli

import (
	"github.com/hermeznetwork/hermez-node/log"
	"os"

	"github.com/hermeznetwork/price-updater-service/config"
	"github.com/spf13/cobra"
)

var (
	cfg     config.Config
	err     error
	rootCmd = &cobra.Command{
		Use: "price-updater",
	}
)

func init() {
	cobra.OnInitialize(baseSetup)

	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(updaterCmd)
	rootCmd.AddCommand(changeProviderCmd)
	rootCmd.AddCommand(setupOriginCmd)
	rootCmd.AddCommand(setupAPIKeyCmd)
	rootCmd.AddCommand(updateStaticTokenCmd)
	rootCmd.AddCommand(priorityCmd)

	serverCmd.Flags().String("pg-user", "", "postgresql username")
	serverCmd.Flags().String("pg-pass", "", "postgresql password")
	serverCmd.Flags().String("pg-host", "", "postgresql host")
	serverCmd.Flags().String("pg-dbname", "", "postgresql database name")
	serverCmd.Flags().Int("pg-port", 5432, "postgresql port")
	serverCmd.Flags().String("eth-network", "", "ethereum address")
	serverCmd.Flags().String("fiat-key", "", "fiat api key")

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}

func baseSetup() {
	cfg, err = config.LoadByEnv()
	if err != nil {
		serverCmd.MarkFlagRequired("pg-user")
		serverCmd.MarkFlagRequired("pg-pass")
		serverCmd.MarkFlagRequired("pg-host")
		serverCmd.MarkFlagRequired("pg-dbname")
		cfg, err = config.LoadByFlags(serverCmd)
		if err != nil {
			panic(err.Error())
		}
	}
}
