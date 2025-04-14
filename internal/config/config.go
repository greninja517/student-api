package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Address string
}

type Config struct {
	Env         string `yaml:"env"`
	StoragePath string `yaml:"storage_path"`
	HttpServer  `yaml:"http_server"`
}

func ConfigurationLoader() *Config {
	var cfg Config

	configfile := os.Getenv("CONFIG_FILE")

	if configfile == "" {
		flag.StringVar(&configfile, "config", "", "configuration file path")
		flag.Parse()

		if configfile == "" {
			log.Fatal("Please provide the config file")
		}
	}

	if _, err := os.Stat(configfile); os.IsNotExist(err) {
		log.Fatalf("%s file doesn't exist", configfile)
	}

	err := cleanenv.ReadConfig(configfile, &cfg)
	if err != nil {
		log.Fatal("Failed to read the configuration..", err)
	}

	return &cfg
}
