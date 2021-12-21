/*
Manages the connection to the Database
*/
package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

/*
Open a connection to the database
*/
func Connect() (*sql.DB, error) {
	config, _ := LoadConfig("database.conf")
	db, err := sql.Open("mysql", config.Db.User+":"+config.Db.Password+"@/"+config.Db.Database)
	if err != nil {
		return nil, err
	}

	// Then validate the DSN data passed
	if err = db.Ping(); err != nil {
		return nil, err
	}

	fmt.Printf("Opening DB ...\n")

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(10)

	return db, nil
}
