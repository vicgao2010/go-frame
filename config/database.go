package config

type Database struct {
	Dns     string `yaml:"dns,omitempty"`
	MaxIdle  int    `yaml:"max-idle,omitempty"`
	MaxOpen  int    `yaml:"max-open,omitempty"`
	TimeOut  string `yaml:"timeout,omitempty"`
	LifeTime string `yaml:"life-time,omitempty"`
	IdleTime string `yaml:"idle-time,omitempty"`
	Prefix   string `yaml:"prefix,omitempty"`
}
