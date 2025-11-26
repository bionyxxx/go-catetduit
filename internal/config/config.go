package config

import (
	"os"
	"strconv"
)

type Config struct {
	IsProduction            bool
	APIPort                 int
	JWTSecret               string
	JWTExpiredInHour        int
	JWTRefreshExpiredInHour int
}

func NewConfig() *Config {
	environtment := os.Getenv("ENVIRONMENT")

	isProduction := false
	if environtment == "production" {
		isProduction = true
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
		IsProduction:            isProduction,
		APIPort:                 apiPort,
		JWTSecret:               jwtSecret,
		JWTExpiredInHour:        jwtExpiredInHour,
		JWTRefreshExpiredInHour: jwtRefreshExpiredInHour,
	}
}
