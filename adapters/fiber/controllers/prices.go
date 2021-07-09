package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hermeznetwork/price-updater-service/core/services"
)

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
	return ctx.JSON(prices)
}
