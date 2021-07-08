package fiber

import (
	"fmt"

	fb "github.com/gofiber/fiber/v2"
	"github.com/hermeznetwork/price-updater-service/config"
)

type Server struct {
	fiber *fb.App

	// controllers
}

func NewServer() *Server {
	server := &Server{
		fiber: fb.New(),
	}
	return server
}

func (s *Server) Start(cfg config.HTTPServerConfig) {
	local := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	s.fiber.Listen(local)
}
