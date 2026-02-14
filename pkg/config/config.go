package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MaplsAccessToken string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	token := os.Getenv("MAPPLS_ACCESS_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("missing required environment variable: MAPPLS_ACCESS_TOKEN")
	}

	return &Config{
		MaplsAccessToken: token,
	}, nil
}
