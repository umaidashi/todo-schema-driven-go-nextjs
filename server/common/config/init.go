package config

import (
	"github.com/spf13/viper"
)

type ConfigType struct {
	DB_DSN string
	PORT   string
}

var Config ConfigType

func Init() {
	viper.SetConfigName(Env.Name())
	viper.SetConfigType("env")
	viper.AddConfigPath("common/config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic("Error loading .env file")
	}

	if err := viper.Unmarshal(&Config); err != nil {
		panic("Error Unmarshal .env file")
	}
}
