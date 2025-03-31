package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port int
}

func LoadConfigs() *Config {
	log.Println("Loading configurations...")

	viper.SetConfigFile("config/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("Failed to load configs:", err.Error())
	}

	c := &Config{
		Port: viper.GetInt("service.port"),
	}

	log.Println(c.GetConfigs())
	log.Println("Configurations LOADED!")

	return c
}

func (c *Config) GetConfigs() string {
	configStr :=
		fmt.Sprintln("Configurations...") +
			fmt.Sprintln("Port:", c.Port)

	return configStr
}
