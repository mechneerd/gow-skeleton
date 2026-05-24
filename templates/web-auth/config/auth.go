package config

import "os"

type AuthConfig struct {
	Defaults struct {
		Guard string
	}
	Guards map[string]GuardConfig
	Providers map[string]ProviderConfig
}

type GuardConfig struct {
	Driver   string
	Provider string
}

type ProviderConfig struct {
	Driver string
	Model  string
}

func LoadAuth() AuthConfig {
	return AuthConfig{
		Defaults: struct {
			Guard string
		}{Guard: getEnv("AUTH_DEFAULT_GUARD", "web")},
		Guards: map[string]GuardConfig{
			"web": {
				Driver:   "session",
				Provider: "users",
			},
			"api": {
				Driver:   "sanctum",
				Provider: "users",
			},
		},
		Providers: map[string]ProviderConfig{
			"users": {
				Driver: "eloquent",
				Model:  "app/Models/User",
			},
		},
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
