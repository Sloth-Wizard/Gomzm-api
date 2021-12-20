package killmails

import (
	"encoding/json"
	"fmt"

	kmm "gomzm-api/models/killmails"
)

func KillmailsList() ([]byte, error) {
	fmt.Printf("Calling API ...\n")
	killmails, err := kmm.GetList()

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
