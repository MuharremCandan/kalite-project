package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	AccessTokenDuration time.Duration `yaml:"accesstokenduration"`
	SecretKey           string        `yaml:"secretkey"`
	HttpServer          struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"httpserver"`

	Database struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
		Name string `yaml:"name"`
	} `yaml:"database"`

	Elastic struct {
		Url string `yaml:"url"`
	} `yaml:"elastic"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
