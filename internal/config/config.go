package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type ApplicationConfig struct {
	Debug          bool          `default:"true"`
	MigrationsPath string        `required:"true" default:"./migrations"`
	TokenTTL       time.Duration `envconfig:"token_ttl" default:"1h"`
}

type GRPCConfig struct {
	Port    int
	Timeout time.Duration
}

type Config struct {
	Application ApplicationConfig
	GRPC        GRPCConfig
}

func MustLoad() *Config {
	var appConfig ApplicationConfig
	err := envconfig.Process("app", &appConfig)
	if err != nil {
		panic(err)
	}

	var grpcConfig GRPCConfig
	err = envconfig.Process("grpc", &grpcConfig)
	if err != nil {
		panic(err)
	}

	cfg := Config{
		Application: appConfig,
		GRPC:        grpcConfig,
	}

	return &cfg
}
