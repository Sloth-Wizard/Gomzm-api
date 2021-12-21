package killmails

import (
	"encoding/json"
	"fmt"

	kmm "gomzm-api/models/killmails"
)

/*
Get a list of killmails from the DB and prepare the data as bytes
*/
func KillmailsList() ([]byte, error) {
	fmt.Printf("Calling API ...\n")

	killmails, err := kmm.Get("list")
	if err != nil {
		return []byte("null"), err
	}

	// Prepare the body
	body, err := json.Marshal(killmails)
	if err != nil {
		return []byte("null"), err
	}

	return body, nil
}
