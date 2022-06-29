package config

type Server struct {
	Addr    string `yaml:"addr,omitempty"`
	Timeout string `yaml:"timeout,omitempty"`
}
