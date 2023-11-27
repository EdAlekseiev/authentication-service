package main

import (
	"github.com/EdAlekseiev/authentication-service/internal/config"
	"github.com/EdAlekseiev/authentication-service/internal/lib/logger"
)

func main() {
	cfg := config.MustLoad()
	logger := logger.NewLogger(cfg)

	logger.Info("Start auth service")
}
