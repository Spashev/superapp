package config

import (
	"os"
)

type Config struct {
	DatabaseDSN string
}

func NewConfig() *Config {
	return &Config{
		DatabaseDSN: os.Getenv("DATABASE_DSN"),
	}
}
