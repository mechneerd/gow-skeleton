package config

import "os"

type AppConfig struct {
	AppName string
	Env     string
	Port    string
}

func Load() AppConfig {
	return AppConfig{
		AppName: getEnv("APP_NAME", "{{ .AppName }}"),
		Env:     getEnv("APP_ENV", "local"),
		Port:    getEnv("APP_PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
