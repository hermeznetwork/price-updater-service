package v1

import (
	"strconv"

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

type TokensController struct {
	svc *services.PriceUpdaterService
}

func NewTokensController(s *services.PriceUpdaterService) *TokensController {
	return &TokensController{
		svc: s,
	}
}

func (p *TokensController) GetTokenPrice(ctx *fiber.Ctx) error {
	// TODO: Getting from memory
	tokenID, err := strconv.Atoi(ctx.Params("token_id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	price, err := p.svc.GetToken(uint(tokenID))
	if price.UsdUpdate == "" {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "token not found",
		})
	}
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.JSON(domainToHttpToken(price))
}

func (p *TokensController) GetTokenPrices(ctx *fiber.Ctx) error {
	// TODO: Getting from memory
	prices, err := p.svc.GetTokens()
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.JSON(fiber.Map{
		"tokens": domainToHttpTokens(prices),
	})
}

func domainToHttpToken(p domain.Token) tokenOut {
	return tokenOut{
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
}

func domainToHttpTokens(prices []domain.Token) []tokenOut {
	var returnTokens []tokenOut

	for _, p := range prices {
		returnTokens = append(returnTokens, domainToHttpToken(p))
	}
	return returnTokens
}
