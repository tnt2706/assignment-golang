package configs

import (
	"log"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type ContainerConfig struct {
	Port      string `env:"PORT" default:"8080"`
	HostName  string `env:"HOST_NAME" default:"assignment"`
	Namespace string `env:"NAMESPACE" default:"alpha"`
	NodeEnv   string `env:"NODE_ENV" default:"development"`
}

func GetContainerConfig() *ContainerConfig {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}

	cfg := ContainerConfig{}

	err = env.Parse(&cfg)
	if err != nil {
		log.Fatalf("unable to parse environment variables: %e", err)
	}

	return &cfg
}
