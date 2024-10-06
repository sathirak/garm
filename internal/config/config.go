package config

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

type AppConfig struct {
	Port     string
	ApiToken string
	Env      string
}

type Config struct {
	App      AppConfig
	Database DatabaseConfig
}
