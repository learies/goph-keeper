package config

import (
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App  AppConfig    `yaml:"app"`
	Log  LoggerConfig `yaml:"logger"`
	GRPC GRPCConfig   `yaml:"grpc"`
	PG   DSNConfig    `yaml:"dsn"`
}

type AppConfig struct {
	SecretKey string `json:"secret_key" env:"APP_SECRET_KEY" env-default:"secret"`
	TokenExp  int    `json:"token_exp" env:"APP_TOKEN_EXP" env-default:"120" env-description:"token expiration (minutes)"`
}

type LoggerConfig struct {
	Level string `json:"log_level" env:"LOG_LEVEL" env-default:"info"`
}

type GRPCConfig struct {
	Address string `json:"grpc_address" env:"GRPC_ADDRESS" env-default:":3200"`
}

type DSNConfig struct {
	DatabaseDSN      string `json:"database_dsn" env:"DATABASE_DSN"`
	MigrationVersion int64  `json:"migration_version" env:"DB_MIGRATION_VERSION" env-default:"1"`
	ConnectTimeout   int    `json:"connect_timeout" env:"DB_CONNECT_TIMEOUT" env-default:"60"`
	QueryTimeout     int    `json:"query_timeout" env:"DB_QUERY_TIMEOUT" env-default:"10"`
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}

func MustLoadConfig() *Config {
	path := fetchConfigPath()

	if path == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}
