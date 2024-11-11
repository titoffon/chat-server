package config

import (
	"github.com/spf13/viper"
)

type Config struct{
	Database struct{
		Host string
		Port int
		User string
		Password string
		DBName string
	}
} 

func LoadConfig() (*Config, error){
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil,err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}