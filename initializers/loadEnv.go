package initializers

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	PORT                    string `env:"PORT"`
	
	DB_HOST                 string `env:"DB_HOST"`
	DB_PORT                 string `env:"DB_PORT"`
	DB_DATABASE             string `env:"DB_DATABASE"`
	DB_USERNAME             string `env:"DB_USERNAME"`
	DB_PASSWORD             string `env:"DB_PASSWORD"`

	TOKEN_LIFESPAN_IN_HOURS int    `env:"TOKEN_LIFESPAN_IN_HOURS"`
	SECRET_TOKEN            string `env:"SECRET_TOKEN"`
}

var Cfg Config

func LoadConfig() () {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Unable to load .env file: %e", err)
	}

	err = env.Parse(&Cfg)
	if err != nil {
		log.Fatalf("Unable to parse env variables: %e", err)
	}
}