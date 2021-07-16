package cli

import (
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
}

func Execute() {
	// TODO: remove panic
	if err := rootCmd.Execute(); err != nil {
		panic(err.Error())
	}
}

func baseSetup() {
	cfg = config.Load()
}
