package logger

import (
	"os"

	"github.com/EdAlekseiev/authentication-service/internal/config"
	"go.uber.org/zap"
)

func NewLogger(cfg *config.Config) *zap.Logger {
	encoderCfg := zap.NewProductionEncoderConfig()
	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       cfg.Application.Debug,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			"stderr",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
		InitialFields: map[string]interface{}{
			"pid": os.Getpid(),
		},
	}

	return zap.Must(config.Build())
}
