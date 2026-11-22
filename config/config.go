package config

import (
	"log"
	"sync"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	App   AppConfig
	Mongo MongoConfig
}

type AppConfig struct {
	AppName         string `envconfig:"APP_NAME" default:"portifyx-market-data-service"`
	AppEnv          string `envconfig:"APP_ENV" default:"local"`
	AppVersion      string `envconfig:"APP_VERSION" default:"1.0.0"`
	AppPort         string `envconfig:"APP_PORT" default:"8083"`
	HTTPTimeoutInMs int64  `envconfig:"HTTP_TIMEOUT_IN_MS" default:"30000"`
	TimeZone        string `envconfig:"TIMEZONE" default:"Asia/Bangkok"`
}

type MongoConfig struct {
	URI          string `envconfig:"MONGO_URI" default:"mongodb://localhost:27017"`
	DatabaseName string `envconfig:"MONGO_DB_NAME" default:"portifyx_marketdata"`
}

var (
	cfg  Config
	once sync.Once
)

func Init() {
	once.Do(func() {
		_ = godotenv.Load()
		if err := envconfig.Process("", &cfg); err != nil {
			log.Fatalf("failed to parse environment config: %v", err)
		}
	})
}

func GetConfig() *Config {
	return &cfg
}
