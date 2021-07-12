package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/hermeznetwork/price-updater-service/adapters/background"
	"github.com/hermeznetwork/price-updater-service/adapters/bitfinex"
	"github.com/hermeznetwork/price-updater-service/adapters/command"
	"github.com/hermeznetwork/price-updater-service/adapters/fiber"
	"github.com/hermeznetwork/price-updater-service/adapters/fiber/controllers"
	"github.com/hermeznetwork/price-updater-service/adapters/postgres"
	"github.com/hermeznetwork/price-updater-service/config"
	"github.com/hermeznetwork/price-updater-service/core/services"
)

func main() {
	Start(config.Load())
}

func Start(cfg config.Config) {
	ctx := context.Background()
	postgresConn := postgres.NewConnection(ctx, &cfg.Postgres)

	// provider
	bitfitexSymbols := []string{"0=ETH", "2=UST", "9=SUSHI", "5=WBT"}
	bitfinexClient := bitfinex.NewClient(bitfitexSymbols)
	// repostitory
	priceRepository := postgres.NewTokenRepository(postgresConn)
	// service
	priceUpdateService := services.NewPriceUpdaterService(bitfinexClient, priceRepository, ctx)
	// command
	cmdUpdatePrice := command.NewUpdatePriceCommand(priceUpdateService)
	// controller
	priceController := controllers.NewPricesController(priceUpdateService)

	server := fiber.NewServer(priceController)

	go func(server *fiber.Server, cfg config.HTTPServerConfig) {
		server.Start(cfg)
	}(server, cfg.HTTPServer)

	bg := background.NewBackground(ctx, cmdUpdatePrice)
	bg.AddWg(1)
	go bg.StartUpdateProcess()
	waitSigInt()
	bg.Stop()
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
