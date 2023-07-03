package configs

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type MongoConfig struct {
	COMMON_DB string `required:"true"`
}

func GetMongoConfig() *MongoConfig {
	var config MongoConfig
	err := envconfig.Process("dynamodb", &config)
	if err != nil {
		log.Fatal(err.Error())
	}

	return &config
}
