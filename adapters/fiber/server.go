package fiber

import (
	"fmt"

	fb "github.com/gofiber/fiber/v2"
	"github.com/hermeznetwork/price-updater-service/adapters/fiber/controllers"
	"github.com/hermeznetwork/price-updater-service/config"
)

type Server struct {
	fiber *fb.App

	*controllers.PricesController
}

func NewServer(pc *controllers.PricesController) *Server {
	server := &Server{
		fiber:            fb.New(),
		PricesController: pc,
	}

	server.fiber.Get("/prices", server.GetPrices)
	server.fiber.Get("/health", func(ctx *fb.Ctx) error {
		return ctx.SendString("Healthy")
	})
	return server
}

func (s *Server) Start(cfg config.HTTPServerConfig) {
	local := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	s.fiber.Listen(local)
}
