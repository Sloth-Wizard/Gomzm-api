package killmails

import (
	"gomzm-api/utils/database"
)

type Killmail struct {
	id                      int
	killmail_id             int
	victim_character_id     int
	final_blow_character_id int
	final_blow_faction_id   int
	killmail_details        string
}

const LIMIT = "2" // Must be a string because it's used in queries

var ( // Database connection init
	db, _ = database.Connect()
)

func KillmailsList() ([]Killmail, error) {
	var killmail Killmail
	var results []Killmail

	// Prepare the statement
	stmtOut, err := db.Prepare("SELECT * FROM killmails LIMIT " + LIMIT)
	if err != nil {
		return nil, err
	}
	defer stmtOut.Close()

	// Execute the statement
	rows, err := stmtOut.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Scan results
	for rows.Next() {
		err := rows.Scan(&killmail.id, &killmail.killmail_id, &killmail.victim_character_id, &killmail.final_blow_character_id, &killmail.final_blow_faction_id, &killmail.killmail_details)
		if err != nil {
			return nil, err
		}
		results = append(results, killmail)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return results, nil
}
