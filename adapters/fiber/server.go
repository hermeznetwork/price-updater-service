package fiber

import (
	"fmt"
	"strings"

	fb "github.com/gofiber/fiber/v2"
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

	server.fiber.Use(func(ctx *fb.Ctx) error {
		origins, err := server.pcr.LoadOriginValue()
		if err != nil {
			return err
		}
		origin := ctx.Get(fb.HeaderOrigin)
		allowOrigin := ""
		allowOrigins := strings.Split(strings.Replace(origins, " ", "", -1), ",")
		for _, allowed := range allowOrigins {
			if allowed == "*" {
				allowOrigin = origin
				break

			}
			if allowed == origin {
				allowOrigin = allowed
				break
			}

			if matchSubdomain(origin, allowed) {
				allowOrigin = origin
				break
			}
		}

		if allowOrigin == "" {
			return ctx.SendStatus(fb.StatusNoContent)
		}
		return ctx.Next()
	})

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
