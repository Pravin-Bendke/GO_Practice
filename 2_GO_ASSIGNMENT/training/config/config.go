package config

import (
	"os"
	"time"
)

type Config struct {
	AppPort     string
	AppEnv      string
	DatabaseURL string
	JWTSecret   string
	JWTExpiry   time.Duration
}

func LocalConfig() Config {
	port := os.Getenv("APP_PORT")
	env := os.Getenv("APP_ENV")
	databaseURL := os.Getenv("DATABASE_URL")
	jwtSecret := os.Getenv("JWT_SECRET")
	jwtExpiry := os.Getenv("JWT_EXPIRY")

	if port == "" {
		port = "8080"
	}
	if env == "" {
		env = "local"
	}

	expiry := 24 * time.Hour

	if jwtExpiry != "" {
		parsedExpiry, err := time.ParseDuration(jwtExpiry)
		if err == nil && parsedExpiry > 0 {
			expiry = parsedExpiry
		}
	}

	return Config{
		AppPort:     port,
		AppEnv:      env,
		DatabaseURL: databaseURL,
		JWTSecret:   jwtSecret,
		JWTExpiry:   expiry,
	}
}
