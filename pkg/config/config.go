package config

import (
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func ImportConfig() *Config {
	viper.SetConfigFile("../config.json")

	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	var config *Config

	err = viper.Unmarshal(&config)

	if err != nil {
		panic(err)
	}

	return config
}

var Module = fx.Provide(ImportConfig)
