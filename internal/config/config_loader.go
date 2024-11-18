package config

import (
	"time"

	"github.com/hotelbear/garm/internal/logger"
	"github.com/joho/godotenv"
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

func Initialize() {
	log := logger.Get()

	if err := ensureRequiredEnvsAreAvailable(); err != nil {
		log.Errorw("startup", "package", "config", "error", err.Error())
		return
	}

	config = Config{
		App: AppConfig{
			Env:        getEnv("ENV", "development"),
			Port:       getEnv("PORT", "8080"),
			ApiToken:   getEnv("X_API_TOKEN", ""),
			JWTExpTime: time.Hour * 24 * 30,
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST"),
			Port:     getEnv("DB_PORT"),
			Username: getEnv("DB_USER"),
			Password: getEnv("DB_PASSWORD"),
			DBName:   getEnv("DB_NAME"),
		},
	}

	log.Infow("startup", "package", "config", "status", "ok")
}

func ensureRequiredEnvsAreAvailable() error {
	log := logger.Get()

	// Load .env file if it exists, but don't error if it doesn't
	_ = godotenv.Load()

	viper.AutomaticEnv()

	for _, env := range requiredEnvs {
		if getEnv(env) == "" {
			log.Errorw("startup", "package", "config:env", env, getEnv(env), "error", "error loading required env")
			return nil
		}
		log.Infow("startup", "package", "config:env", env, getEnv(env), "status", "ok")
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
