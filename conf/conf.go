package conf

import (
	"encoding/json"
	"log"
	"os"
)

type DatabaseConfig struct {
	Host	string
	Port	string
	User	string
	Name	string
	Pass	string
}

type Config struct {
	Database	DatabaseConfig
}

func GetConfig(isProd bool) (*Config, error) {
	var conf Config
	var envCOnf string

	if isProd {
		envCOnf = "conf/conf.prod.json"
	} else {
		envCOnf = "conf/conf.dev.json"
	}

	configFile, err := os.Open(envCOnf)

	defer func() {
		if err := configFile.Close(); err != nil {
			log.Fatal("Error::Can't close config file")
		}
	}()

	if err != nil {
		return nil, err
	}

	jsonParser := json.NewDecoder(configFile)

	if err := jsonParser.Decode(&conf); err != nil {
		log.Fatal(err)
	}

	return &Config{conf.Database}, nil
}
