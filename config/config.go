package config

import (
	"errors"

	"github.com/spf13/viper"
)

type Config struct {
	Host  string `mapstructure:"HOST"`
	Port  string `mapstructure:"PORT"`
	DBURL string `mapstructure:"DATABASE_URL"`
}

const (
	defaultPort = "3000"
	defaultHost = "localhost"
)

func GetConfig() (*Config, error) {
	envVar := viper.New()
	envVar.AddConfigPath(".")
	envVar.SetConfigFile(".env")

	cfg := &Config{
		Host: defaultHost,
		Port: defaultPort,
	}

	err := envVar.ReadInConfig()
	if err != nil {
		return nil, err
	}

	if envVar.Get("DATABASE_URL") == "" || envVar.Get("DATABASE_URL") == nil {
		return nil, errors.New("DATABASE_URL is not set")
	}

	err = envVar.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
