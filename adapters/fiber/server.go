package fiber

import (
	"fmt"

	fb "github.com/gofiber/fiber/v2"
	v1 "github.com/hermeznetwork/price-updater-service/adapters/fiber/controllers/v1"
	"github.com/hermeznetwork/price-updater-service/config"
)

type Server struct {
	fiber *fb.App

	*v1.CurrencyController
	*v1.TokensController
}

func NewServer(cc *v1.CurrencyController, tc *v1.TokensController) *Server {
	server := &Server{
		fiber:              fb.New(),
		CurrencyController: cc,
		TokensController:   tc,
	}

	server.fiber.Get("v1/tokens", server.GetTokenPrices)
	server.fiber.Get("v1/tokens/:token_id", server.GetTokenPrice)
	server.fiber.Get("v1/currencies", server.GetFiatPrices)
	server.fiber.Get("v1/currencies/:currency", server.GetFiatPrice)
	server.fiber.Get("v1/health", func(ctx *fb.Ctx) error {
		return ctx.SendString("Healthy")
	})
	return server
}

func (s *Server) Start(cfg config.HTTPServerConfig) {
	local := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	s.fiber.Listen(local)
}
