package config

import (
	"os"
)

type Config struct {
	Env      string
	HTTPAddr string
}

func Load() Config {
	return Config{
		Env:      getEnv("APP_ENV", "development"),
		HTTPAddr: getEnv("HTTP_ADDR", ":8080"),
	}
}

func getEnv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
