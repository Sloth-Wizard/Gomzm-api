package database

import (
	"encoding/json"
	"os"
)

type Config struct {
	Db struct {
		Host     string `json:"Host"`
		User     string `json:"User"`
		Password string `json:"Password"`
		Database string `json:"Database"`
	}
	Server struct {
		Address string `json:"Address"`
		Port    string `json:"Port"`
	}
}

/*
Load the database configuration file
*/
func LoadConfig(filename string) (*Config, error) {
	var config Config
	configFile, err := os.Open("config/" + filename)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)

	if err != nil {
		return nil, err
	}

	return &config, nil
}
