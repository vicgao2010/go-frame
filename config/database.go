package config

type Database struct {
	Dns      string `yaml:"dns"`
	MaxIdle  int    `yaml:"maxIdle"`
	MaxOpen  int    `yaml:"maxOpen"`
	TimeOut  string `yaml:"timeout"`
	LifeTime string `yaml:"lifeTime"`
	IdleTime string `yaml:"idleTime"`
	Prefix   string `yaml:"prefix"`
}
