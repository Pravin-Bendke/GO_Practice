package config

import "os"

type Config struct {
	AppPort     string
	AppEnv      string
	DatabaseURL string
}

func LocalConfig() Config {
	port := os.Getenv("APP_PORT")
	env := os.Getenv("APP_ENV")
	databaseURL := os.Getenv("DatabaseURL")

	if port == "" {
		port = "8080"
	}
	if env == "" {
		env = "local"
	}

	return Config{
		AppPort:     port,
		AppEnv:      env,
		DatabaseURL: databaseURL,
	}
}
