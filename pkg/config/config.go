package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Riot struct {
		ApiKey string `yaml:"apikey"`
	} `yaml:"riot"`
}

func ReadConfig() Config {
	yamlFile, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %s", err)
		panic(err)
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Error parsing YAML file: %s", err)
		panic(err)
	}

	return config
}
