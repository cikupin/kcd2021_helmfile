package bootstrap

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

// ConfigObjects defines object for configuration items
type ConfigObjects struct {
	AppOptions AppOptions
	DBOptions  DBOption
}

type AppOptions struct {
	AppPort int `env:"APP_PORT"`
}

func LoadConfig() (cfg ConfigObjects) {
	godotenv.Load("./params/.env")

	err := env.Parse(&cfg)
	if err != nil {
		log.Fatalf("[ERROR] Failed to parse env variable to object: %s\n", err.Error())
	}

	return cfg
}
