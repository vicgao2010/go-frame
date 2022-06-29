package yaml

import (
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

func New(path string) (*viper.Viper, error) {
	Viper := viper.New()
	ext := strings.Replace(filepath.Ext(path), ".", "", 1)
	name := strings.Replace(filepath.Base(path), filepath.Ext(path), "", 1)
	Viper.SetConfigName(name)
	Viper.SetConfigType(ext)
	Viper.AddConfigPath(filepath.Dir(path))
	err := Viper.ReadInConfig()
	return Viper, err
}
