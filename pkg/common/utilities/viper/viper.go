package viper

import (
	"github.com/spf13/viper"
)

//region public functions

// NewViper - it extracts and returns new viper instance
func NewViper() (*viper.Viper, error) {
	var viperInstance = viper.New()
	viperInstance.AddConfigPath(".")
	viperInstance.SetConfigName("config")
	viperInstance.SetConfigType("json")
	if err := viperInstance.ReadInConfig(); err != nil {
		return viperInstance, err
	}
	return viperInstance, nil
}

//endregion
