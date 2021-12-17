package main

import (
	"fmt"
	"gomzm-api/utils/database"
)

func main() {
	config, _ := database.LoadConfig("database.conf")
	fmt.Println(config)
}
