package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func New(file *viper.Viper) *Config {
	var cfg *Config
	if err := file.Unmarshal(&cfg); err != nil {
		panic(fmt.Sprintf("err:%s", err))
	}
	return cfg
}

type Config struct {
	App      App      `yaml:"app,omitempty"`
	Logger   Logger   `yaml:"logger,omitempty"`
	Database Database `yaml:"logger,omitempty"`
	Redis    Redis    `yaml:"redis,omitempty"`
	Server   Server   `yaml:"server,omitempty"`
}
