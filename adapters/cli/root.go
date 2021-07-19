package cli

import (
	"log"
	"os"

	"github.com/hermeznetwork/price-updater-service/config"
	"github.com/spf13/cobra"
)

var (
	cfg     config.Config
	rootCmd = &cobra.Command{
		Use: "price-updater",
	}
)

func init() {
	cobra.OnInitialize(baseSetup)

	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(updaterCmd)
	rootCmd.AddCommand(changeProviderCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}

func baseSetup() {
	cfg = config.Load()
}
