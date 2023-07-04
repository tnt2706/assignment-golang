package configs

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type MongoConfig struct {
	DbCommonConnectString string `env:"DB_COMMON_CONNECT_STRING" required:"true"`
}

func GetMongoConfig() *MongoConfig {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}

	cfg := MongoConfig{}

	err = env.Parse(&cfg)
	if err != nil {
		log.Fatalf("unable to parse environment variables: %e", err)
	}
	fmt.Printf("%+v\n", cfg)
	return &cfg
}
