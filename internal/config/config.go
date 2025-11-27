package config

import (
	"os"
	"strconv"
)

type Config struct {
	IsProduction            bool
	Domain                  string
	APIPort                 int
	JWTSecret               string
	JWTExpiredInHour        int
	JWTRefreshExpiredInHour int
}

func NewConfig() *Config {
	environment := os.Getenv("APP_ENVIRONMENT")

	isProduction := false
	if environment == "production" {
		isProduction = true
	}

	domain := os.Getenv("APP_DOMAIN")

	if domain == "" {
		panic("APP_DOMAIN environment variable is required")
	}

	apiPort, err := strconv.Atoi(os.Getenv("API_PORT"))

	if err != nil {
		apiPort = 8080 // default port
	}

	jwtSecret := os.Getenv("JWT_SECRET")

	if jwtSecret == "" {
		panic("JWT_SECRET environment variable is required")
	}

	jwtExpiredInHour, err := strconv.Atoi(os.Getenv("JWT_EXPIRED_IN_HOURS"))

	if err != nil || jwtExpiredInHour <= 0 {
		jwtExpiredInHour = 24 // default to 24 hours
	}

	jwtRefreshExpiredInHour, err := strconv.Atoi(os.Getenv("JWT_REFRESH_EXPIRED_IN_HOURS"))

	if err != nil || jwtRefreshExpiredInHour <= 0 {
		jwtRefreshExpiredInHour = 168 // default to 7 days
	}

	return &Config{
		Domain:                  domain,
		IsProduction:            isProduction,
		APIPort:                 apiPort,
		JWTSecret:               jwtSecret,
		JWTExpiredInHour:        jwtExpiredInHour,
		JWTRefreshExpiredInHour: jwtRefreshExpiredInHour,
	}
}
