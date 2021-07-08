package main

import (
	"context"
	"fmt"

	"github.com/hermeznetwork/price-updater-service/adapters/bitfinex"
	"github.com/hermeznetwork/price-updater-service/adapters/coingecko"
	"github.com/hermeznetwork/price-updater-service/config"
)

func main() {
	Start(config.Load())
}

func Start(cfg config.Config) {
	ctx := context.Background()

	/*
		// setup of running provider
		bitfinexClient := bitfinex.NewClient()
		coingeckoClient := coingecko.NewClient()

		// TODO: setup with database
		mongodb := mongo.NewClient(ctx, &cfg)
		// TODO: setup with repository (pass the database connection)
		// TODO: setup with services

		priceProviders := []ports.PriceProvider{
			bitfinexClient, coingeckoClient,
		}
		providerSelectorService := services.NewProviderSelectorService(priceProviders)
		selectedProvider := providerSelectorService.Select("coingecko")

		priceUpdaterService := services.NewPriceUpdaterService(selectedProvider)
		priceUpdaterService.UpdatePrices()
		server := fiber.NewServer()
		server.Start(cfg.HTTPServer)
	*/

	client := coingecko.NewClient()
	client = bitfinex.NewClient()
	r, e := client.GetPrices(ctx)
	fmt.Println("error: ", e)
	fmt.Println("result: ", r)
}
