package config

type Redis struct {
	Server   string `yaml:"server,omitempty"`
	PassWord string `yaml:"password,omitempty"`
	DataBase int    `yaml:"database,omitempty"`
	PoolSize int    `yaml:"pool-size,omitempty"`
}
