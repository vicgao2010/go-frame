package config

type Server struct {
	Addr    string `yaml:"addr"`
	Timeout string `yaml:"timeout"`
}
