package config

type Redis struct {
	Server   string `yaml:"server"`
	PassWord string `yaml:"password"`
	DataBase int    `yaml:"database"`
	PoolSize int    `yaml:"poolSize"`
}
