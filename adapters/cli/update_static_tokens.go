package cli

import (
	"context"
	"github.com/hermeznetwork/hermez-node/log"
	"os"

	"github.com/hermeznetwork/price-updater-service/adapters/postgres"
	"github.com/hermeznetwork/price-updater-service/config"
	"github.com/hermeznetwork/price-updater-service/core/services"
	"github.com/spf13/cobra"
)

var staticTokenID uint
var newStaticTokenPrice float64

var updateStaticTokenCmd = &cobra.Command{
	Use:   "update-static-token",
	Short: "command to update a static token price",
	Run: func(cmd *cobra.Command, args []string) {
		updateStaticToken(cfg)
	},
}

func init() {
	updateStaticTokenCmd.Flags().UintVar(&staticTokenID, "tokenID", 0, "static token id")
	updateStaticTokenCmd.Flags().Float64Var(&newStaticTokenPrice, "price", 0.0, "static token price")
	updateStaticTokenCmd.MarkFlagRequired("tokenID")
	updateStaticTokenCmd.MarkFlagRequired("price")
}

func updateStaticToken(cfg config.Config) {
	log.Info("update static token %d\n", staticTokenID)
	ctx := context.Background()
	postgresCon := postgres.NewConnection(ctx, &cfg.Postgres)
	defer postgresCon.DB.Close()
	tokenRepository := postgres.NewTokenRepository(postgresCon)
	tokenService := services.NewPriceUpdaterService(ctx, nil, tokenRepository)
	err := tokenService.UpdatePrice(staticTokenID, newStaticTokenPrice)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	log.Info("static token value updated!")
}
