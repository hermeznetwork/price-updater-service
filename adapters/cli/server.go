package cli

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/hermeznetwork/price-updater-service/adapters/background"
	"github.com/hermeznetwork/price-updater-service/adapters/bbolt"
	"github.com/hermeznetwork/price-updater-service/adapters/command"
	"github.com/hermeznetwork/price-updater-service/adapters/fiat"
	"github.com/hermeznetwork/price-updater-service/adapters/fiber"
	v1 "github.com/hermeznetwork/price-updater-service/adapters/fiber/controllers/v1"
	"github.com/hermeznetwork/price-updater-service/adapters/postgres"
	"github.com/hermeznetwork/price-updater-service/config"
	"github.com/hermeznetwork/price-updater-service/core/services"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		server(cfg)
	},
}

func server(cfg config.Config) {
	ctx := context.Background()
	postgresConn := postgres.NewConnection(ctx, &cfg.Postgres)
	bboltConn := bbolt.NewConnection(cfg.Bbolt)
	configProviderRepo := bbolt.NewConfigProviderRepository(bboltConn)

	priceSelector := services.NewProviderSelectorService(configProviderRepo, cfg)

	// providers
	fiatProvider := fiat.NewClient(cfg.Fiat.APIKey)
	tokenProvider, err := priceSelector.CurrentProvider()
	if err != nil {
		log.Println("try server start up:", err.Error())
		os.Exit(1)
	}

	// repostitory
	priceRepository := postgres.NewTokenRepository(postgresConn)
	fiatRepository := postgres.NewFiatPricesRepository(postgresConn)

	// service
	tokenPriceUpdateService := services.NewPriceUpdaterService(ctx, tokenProvider, priceRepository)
	fiatPriceUpdateService := services.NewFiatUpdaterServices(fiatRepository, fiatProvider)

	// command
	cmdUpdatePrice := command.NewUpdatePriceCommand(tokenPriceUpdateService)
	cmdUpdateFiatePrice := command.NewUpdateFiatPriceCommand(fiatPriceUpdateService)

	// controllers
	tokenController := v1.NewTokensController(tokenPriceUpdateService)
	currencyController := v1.NewCurrencyController(fiatPriceUpdateService)

	server := fiber.NewServer(currencyController, tokenController)

	go func(server *fiber.Server, cfg config.HTTPServerConfig) {
		server.Start(cfg)
	}(server, cfg.HTTPServer)

	bgFiat := background.NewBackground(ctx, cmdUpdateFiatePrice)
	bgToken := background.NewBackground(ctx, cmdUpdatePrice)
	bgFiat.AddWg(1)
	bgToken.AddWg(1)
	go bgToken.StartUpdateProcess()
	go bgFiat.StartUpdateProcess()
	waitSigInt()
	bgToken.Stop()
	bgFiat.Stop()
}

func waitSigInt() {
	stopCh := make(chan interface{})

	ossig := make(chan os.Signal, 1)
	signal.Notify(ossig, os.Interrupt)

	go func() {
		for sig := range ossig {
			if sig == os.Interrupt {
				stopCh <- nil
			}
		}
	}()
	<-stopCh
}
