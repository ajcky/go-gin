package config

import (
	"github.com/spf13/viper"
)

var Cfg *viper.Viper

func InitConfig() error {
	configViper := viper.New()
	configViper.SetConfigName("default")
	configViper.SetConfigType("yml")
	configViper.AddConfigPath("/app/env")
	err := configViper.ReadInConfig()
	if err != nil {
		return err
	}
	Cfg = configViper
	return nil

}
