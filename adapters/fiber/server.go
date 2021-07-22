package fiber

import (
	"fmt"
	"strings"

	fb "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	v1 "github.com/hermeznetwork/price-updater-service/adapters/fiber/controllers/v1"
	"github.com/hermeznetwork/price-updater-service/config"
	"github.com/hermeznetwork/price-updater-service/core/ports"
)

type Server struct {
	fiber *fb.App
	db    ports.MemoryRepository

	*v1.CurrencyController
	*v1.TokensController
}

func NewServer(cc *v1.CurrencyController, tc *v1.TokensController, db ports.MemoryRepository) *Server {
	server := &Server{
		fiber:              fb.New(),
		CurrencyController: cc,
		TokensController:   tc,
	}

	server.fiber.Use(cors.New(cors.Config{
		Next: func(ctx *fb.Ctx) bool {
			origins, err := db.GetAllowOriginValues()
			if err != nil {
				return false
			}
			origin := ctx.Get(fb.HeaderOrigin)
			allowOrigins := strings.Split(strings.Replace(origins, " ", "", -1), ",")
			for _, allowed := range allowOrigins {
				if allowed == "*" {
					return true
				}
				if allowed == origin {
					return true
				}
			}
			return false
		},
	}))

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
