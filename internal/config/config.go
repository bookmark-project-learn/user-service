package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Config struct for app config
type Config struct {
	AppPort  string `envconfig:"APP_PORT" default:"8080"`
	BasePath string `envconfig:"BASE_PATH" default:"/"`
}

func LoadConfig() (*Config, error) {
	var cfg Config

	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	return &cfg, nil
}
