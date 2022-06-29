package config

import (
	"fmt"
	"github.com/spf13/viper"
)

const (
	EnvLocal       = "local"
	EnvProduction  = "production"
	EnvDevelopment = "development"
)

func New(file *viper.Viper) *Config {
	var cfg *Config
	if err := file.Unmarshal(&cfg); err != nil {
		fmt.Printf("err:%s", err)
	}
	return cfg
}

type Config struct {
	App      App      `yaml:"app,omitempty"`
	Logger   Logger   `yaml:"logger,omitempty"`
	Database Database `yaml:"logger,omitempty"`
	Redis    Redis    `yaml:"redis,omitempty"`
}
