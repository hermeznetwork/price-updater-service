package v1

import (
	"fmt"
	"strconv"

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

func (p *CurrencyController) GetFiatPrice(ctx *fiber.Ctx) error {
	currency := ctx.Params("currency")
	price, err := p.svc.GetPrice(currency)
	if price.BaseCurrency == "" {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fmt.Sprintf("currency %s not found", currency),
		})
	}
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.JSON(domainToHttpCurrency(price))
}

func (p *CurrencyController) GetFiatPrices(ctx *fiber.Ctx) error {
	// TODO: Getting from memory
	fromItem, err := strconv.Atoi(ctx.Query("fromItem"))
	if err != nil {
		fromItem = 0
	}
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 0
	}
	var order string = ctx.Query("order")
	if order == "" {
		order = "ASC"
	}
	prices, err := p.svc.GetPrices(uint(limit), uint(fromItem), order)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.JSON(fiber.Map{
		"currencies": domainToHttpCurrencies(prices),
	})
}

func domainToHttpCurrency(p domain.FiatPrice) fiatOut {
	return fiatOut{
		Currency:     p.Currency,
		BaseCurrency: p.BaseCurrency,
		Price:        p.Price,
		LastUpdate:   p.LastUpdate,
	}
}

func domainToHttpCurrencies(prices []domain.FiatPrice) []fiatOut {
	var returnTokens []fiatOut

	for _, p := range prices {
		returnTokens = append(returnTokens, domainToHttpCurrency(p))
	}
	return returnTokens
}
