package database

import (
	"encoding/json" // fmt implements formatted I/O.
	"os"
	//_ "github.com/go-sql-driver/mysql"
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
func LoadConfig(filename string) (Config, error) {
	var config Config
	configFile, err := os.Open("config/" + filename)
	if err != nil {
		return config, err
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)

	return config, err
}
