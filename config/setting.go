package config

import (
	"log"

	"github.com/spf13/viper"
)

type Setting struct {
	Port        string `mapstructure:"PORT"`
	DatabaseUrl string `mapstructure:"DATABASE_URL"`
}

// NewSetting - Load env
func NewSetting() Setting {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("cannot load config: %v", err)
	}

	var env Setting
	viper.Unmarshal(&env)
	return env
}
