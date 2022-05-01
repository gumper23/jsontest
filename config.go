package main

import "github.com/spf13/viper"

type Configuration struct {
	API      API      `json:"api"`
	Database Database `json:"database"`
}

func ReadConfig() (config *Configuration, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	if err = viper.ReadInConfig(); err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
