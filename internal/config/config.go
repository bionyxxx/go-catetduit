package config

import (
	"os"
	"strconv"
)

type Config struct {
	APIPort   int
	JWTSecret string
}

func NewConfig() *Config {
	apiPort, err := strconv.Atoi(os.Getenv("API_PORT"))

	if err != nil {
		apiPort = 8080 // default port
	}

	jwtSecret := os.Getenv("JWT_SECRET")

	if jwtSecret == "" {
		panic("JWT_SECRET environment variable is required")
	}

	return &Config{
		APIPort:   apiPort,
		JWTSecret: jwtSecret,
	}
}
