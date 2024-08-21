package config

import (
	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func Load() *Config {

	err := godotenv.Load()

	if err != nil {
		log.Fatal().Msgf("Error loading .env file: %s", err.Error())
	}

	var c Config
	if err := env.Parse(&c); err != nil{
		log.Fatal().Msgf("Unable to parse env: %s", err.Error())
	}

	return &c
}

type Config struct {
	Database Database
}

type Database struct {
	Host     string `env:"DATABASE_HOST"`
	Port     int `env:"DATABASE_PORT"`
	User     string `env:"DATABASE_USER"`
	Password string `env:"DATABASE_PASSWORD"`
	Name   string `env:"DATABASE_NAME"`
}