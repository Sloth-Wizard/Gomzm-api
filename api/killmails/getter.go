package killmails

import (
	"encoding/json"
	"fmt"
	"time"

	"gomzm-api/models/killmails"
)

/*
Get a list of killmails from the DB and prepare the data as bytes
*/
func KillmailsList() ([]byte, error) {
	fmt.Printf("[%s] Calling API ...\n", time.Now().Format("2006-01-02 15:04:05"))

	response, err := killmails.Get("list")
	if err != nil {
		return []byte("null"), err
	}

	// Prepare the body
	body, err := json.Marshal(response)
	if err != nil {
		return []byte("null"), err
	}

	return body, nil
}
