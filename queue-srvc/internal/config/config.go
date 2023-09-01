package config

import (
	"github.com/spf13/viper"
	"log/slog"
	"sync"
)

type Config struct {
	Port string `yaml:"port" mapstructure:"PORT"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	instance = &Config{}
	once.Do(func() {
		viper.AddConfigPath("../app/internal/config/envs")
		viper.SetConfigName("dev")
		viper.SetConfigType("env")

		viper.AutomaticEnv()

		err := viper.ReadInConfig()

		if err != nil {
			slog.Error(err.Error())
		}

		err = viper.Unmarshal(&instance)
	})
	return instance
}
