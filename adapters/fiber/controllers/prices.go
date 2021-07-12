package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hermeznetwork/price-updater-service/core/domain"
	"github.com/hermeznetwork/price-updater-service/core/services"
)

type tokenOut struct {
	ItemID           uint    `json:"itemId"`
	ID               uint    `json:"id"`
	EthereumBlockNum uint    `json:"ethereumBlockNum"`
	EthereumAddress  string  `json:"ethereumAddress"`
	Name             string  `json:"name"`
	Symbol           string  `json:"symbol"`
	Decimals         uint    `json:"decimals"`
	Usd              float64 `json:"USD"`
	UsdUpdate        string  `json:"usdUpdate"`
}

type PricesController struct {
	svc *services.PriceUpdaterService
}

func NewPricesController(s *services.PriceUpdaterService) *PricesController {
	return &PricesController{
		svc: s,
	}
}

func (p *PricesController) GetPrices(ctx *fiber.Ctx) error {
	// TODO: Getting from memory
	prices, err := p.svc.GetTokens()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return ctx.JSON(fiber.Map{
		"tokens": domainToHttpReturn(prices),
	})
}

func domainToHttpReturn(prices []domain.Token) []tokenOut {
	var returnTokens []tokenOut

	for _, p := range prices {
		nTkOut := tokenOut{
			ItemID:           p.ItemID,
			ID:               p.ID,
			EthereumBlockNum: p.BlockNum,
			EthereumAddress:  p.Address,
			Name:             p.Name,
			Symbol:           p.Symbol,
			Decimals:         p.Decimals,
			Usd:              p.Price,
			UsdUpdate:        p.UsdUpdate,
		}
		returnTokens = append(returnTokens, nTkOut)
	}
	return returnTokens
}
