package main

import (
	"fmt"

	"github.com/spf13/viper"
)

// LoadConfig is to load the configration file and load it.
func LoadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}
