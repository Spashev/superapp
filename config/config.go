package config

import (
	"os"
)

type Config struct {
	DatabaseDSN  string
	JWTSecretKey string
}

func NewConfig() *Config {
	return &Config{
		JWTSecretKey: os.Getenv("JWT_SECRET_KEY"),
	}
}
