package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/sathirak/garm/pkg/logger"
	"github.com/spf13/viper"

)

var config Config

var requiredEnvs = [...]string{
	"ENV",
	"PORT",
	"X_API_TOKEN",
	"DB_HOST",
	"DB_PORT",
	"DB_USER",
	"DB_PASSWORD",
	"DB_NAME",
}

func Get() *Config {
	return &config
}

func Initialize() (*Config, error) {
	log := logger.Get()

	if err := ensureRequiredEnvsAreAvailable(); err != nil {
		return nil, err
	}

	config = Config{
		App: AppConfig{
			Env:      getEnv("ENV", "development"),
			Port:     getEnv("PORT", "8080"),
			ApiToken: getEnv("X_API_TOKEN", ""),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST"),
			Port:     getEnv("DB_PORT"),
			Username: getEnv("DB_USER"),
			Password: getEnv("DB_PASSWORD"),
			DBName:   getEnv("DB_NAME"),
		},
	}

	log.Info("config loaded successfully!")
	return &config, nil
}

func ensureRequiredEnvsAreAvailable() error {
	log := logger.Get()
	err := godotenv.Load()
	if err != nil {
		return err
	}

	viper.AutomaticEnv()

	for _, env := range requiredEnvs {
		if getEnv(env) == "" {
			return fmt.Errorf("error loading .env var %s", env)
		} else {
			log.Infow("env var", env, getEnv(env))
		}
	}
	return nil
}

func getEnv(key string, defaultVal ...string) string {
	if value := viper.GetString(key); value != "" {
		return value
	}
	if len(defaultVal) > 0 {
		return defaultVal[0]
	}
	return ""
}

// func getEnvNumber(key string, defaultVal int) int {
// 	if value := viper.GetString(key); value != "" {
// 		if intValue, err := strconv.Atoi(value); err == nil {
// 			return intValue
// 		}
// 	}
// 	return defaultVal
// }
