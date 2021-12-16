package main

import (
	"encoding/json" // fmt implements formatted I/O.
	"fmt"
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
func loadConfig(filename string) (Config, error) {
	var config Config
	configFile, errConfig := os.Open("Config/" + filename)
	defer configFile.Close()

	// Error for the conf file
	if errConfig != nil {
		return config, errConfig
	}

	jsonParser := json.NewDecoder(configFile)
	err := jsonParser.Decode(&config)

	return config, err
}

func main() {
	config, _ := loadConfig("database.conf")
	fmt.Println(config)
}
