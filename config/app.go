package config

type App struct {
	Env    string `yaml:"env,omitempty"`
	Name   string `yaml:"name,omitempty"`
	Debug  bool   `yaml:"debug,omitempty"`
}
