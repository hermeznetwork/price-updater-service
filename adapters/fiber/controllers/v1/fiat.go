package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hermeznetwork/price-updater-service/core/domain"
	"github.com/hermeznetwork/price-updater-service/core/services"
)

type fiatOut struct {
	Currency     string  `json:"currency"`
	BaseCurrency string  `json:"baseCurrency"`
	Price        float64 `json:"price"`
	LastUpdate   string  `json:"lastUpdate"`
}

type CurrencyController struct {
	svc *services.FiatUpdaterService
}

func NewCurrencyController(s *services.FiatUpdaterService) *CurrencyController {
	return &CurrencyController{
		svc: s,
	}
}

func (p *CurrencyController) GetFiatPrices(ctx *fiber.Ctx) error {
	// TODO: Getting from memory
	prices, err := p.svc.GetPrices()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return ctx.JSON(fiber.Map{
		"currencies": domainToHttpCurrencies(prices),
	})
}

func domainToHttpCurrencies(prices []domain.FiatPrice) []fiatOut {
	var returnTokens []fiatOut

	for _, p := range prices {
		nTkOut := fiatOut{
			Currency:     p.Currency,
			BaseCurrency: p.BaseCurrency,
			Price:        p.Price,
			LastUpdate:   p.LastUpdate,
		}
		returnTokens = append(returnTokens, nTkOut)
	}
	return returnTokens
}
