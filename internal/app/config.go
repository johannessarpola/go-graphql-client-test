package app

import (
	"gopkg.in/yaml.v3"
	"os"
)

type API struct {
	Key     string `yaml:"key"`
	Address string `yaml:"address"`
}

type Config struct {
	API  API    `yaml:"api"`
	Port string `yaml:"port"`
}

func LoadConfig[T interface{}](filename string) (T, error) {
	var config T

	configFile, err := os.ReadFile(filename)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
