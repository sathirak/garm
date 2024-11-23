package config

import "time"

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

type AppConfig struct {
	ApiVersion string
	Port       string
	ApiToken   string
	Env        string
	JWTExpTime time.Duration
}

type SMTPConfig struct {
	Username string
	Password string
	Host     string
	Port     string
}

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	SMTP     SMTPConfig
}
