package config

type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

func LoadDatabase() DatabaseConfig {
	return DatabaseConfig{
		Driver:   getEnv("DB_CONNECTION", "{{ .DatabaseDriver }}"),
		Host:     getEnv("DB_HOST", "127.0.0.1"),
		Port:     getEnv("DB_PORT", "3306"),
		Database: getEnv("DB_DATABASE", "{{ .AppName }}"),
		Username: getEnv("DB_USERNAME", "root"),
		Password: getEnv("DB_PASSWORD", ""),
	}
}
