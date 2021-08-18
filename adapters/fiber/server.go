package fiber

import (
	"fmt"

	fb "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/skip"
	v1 "github.com/hermeznetwork/price-updater-service/adapters/fiber/controllers/v1"
	"github.com/hermeznetwork/price-updater-service/config"
	"github.com/hermeznetwork/price-updater-service/core/services"
)

type Server struct {
	fiber *fb.App
	pcr   *services.ProjectConfigService

	*v1.CurrencyController
	*v1.TokensController
}

func NewServer(cc *v1.CurrencyController, tc *v1.TokensController, pcr *services.ProjectConfigService) *Server {
	server := &Server{
		fiber:              fb.New(),
		pcr:                pcr,
		CurrencyController: cc,
		TokensController:   tc,
	}

	allowOrigins, err := server.pcr.LoadOriginValue()
	if err != nil {
		allowOrigins = "*"
	}

	apiKey, err := server.pcr.LoadAPIKey()
	if err != nil {
		apiKey = "pr1c3upd4t3r"
	}

	server.fiber.Use(func(c *fb.Ctx) error {
		// allow CORS requests
		if c.Method() == "OPTIONS" {
			return c.Next()
		}
		if c.Get("X-Api-Key") == apiKey {
			return c.Next()
		}

		return c.SendStatus(fb.StatusNoContent)
	})

	server.fiber.Use(cors.New(cors.Config{AllowOrigins: allowOrigins}))

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
