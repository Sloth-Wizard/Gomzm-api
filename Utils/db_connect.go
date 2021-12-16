package utils

import (
	//"database/sql"
	"encoding/json"
	//"errors"
	"fmt" // fmt implements formatted I/O.
	"os"
	//"strings"
	//"time"
	//_ "github.com/go-sql-driver/mysql"
)

type Configuration struct {
	Databases []string
}

func main() {
	fmt.Println("DB Connect")
}

/*
	Establish a connexion to the database
*/
func connect_db() {

}

/*
	Get the database configuration file
*/
func get_config() (idk???) {
	file, _ := os.Open("Config/conf.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration := Configuration{}

	err := decoder.Decode(&configuration)
	if err != nil {
		errMsg := fmt.Sprintf("Error: %s", err)
	}

	return json_thing
}
