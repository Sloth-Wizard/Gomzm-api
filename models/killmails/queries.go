/*
Contains all the queries to get killmails from the database
*/
package killmails

import (
	"encoding/json"
	"fmt"
	"gomzm-api/utils/database"
)

type Killmail struct {
	Id                      uint            `json:"id"`
	Killmail_id             uint            `json:"killmail_id"`
	Victim_character_id     uint            `json:"victim_character_id"`
	Final_blow_character_id uint            `json:"final_blow_character_id"`
	Final_blow_faction_id   uint            `json:"final_blow_faction_id"`
	Killmail_details        json.RawMessage `json:"killmail_details"`
}

const LIMIT = "40" // Must be a string because it's used in queries

func GetList() ([]Killmail, error) {
	var killmail Killmail
	var results []Killmail

	// Open DB connection
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	// Prepare the statement
	stmtOut, err := db.Prepare("SELECT * FROM killmails LIMIT " + LIMIT)
	if err != nil {
		return nil, err
	}
	defer stmtOut.Close()

	// Execute the statement
	rows, err := stmtOut.Query()
	fmt.Printf("Getting KM from DB ...\n")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	defer db.Close()
	fmt.Printf("Closing DB ...\n")

	// Scan results
	for rows.Next() {
		err := rows.Scan(
			&killmail.Id,
			&killmail.Killmail_id,
			&killmail.Victim_character_id,
			&killmail.Final_blow_character_id,
			&killmail.Final_blow_faction_id,
			&killmail.Killmail_details)

		if err != nil {
			return nil, err
		}

		results = append(results, killmail)
	}

	return results, nil
}
