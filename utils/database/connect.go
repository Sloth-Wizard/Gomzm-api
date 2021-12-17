package database

import (
	"database/sql"
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
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(10)

	return db, nil
}
