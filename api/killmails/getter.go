package killmails

import (
	"encoding/json"
	kmm "gomzm-api/models/killmails"
)

func KillmailsList() ([]byte, error) {
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
