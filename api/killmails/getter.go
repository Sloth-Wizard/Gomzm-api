package killmails

import (
	kmm "gomzm-api/models/killmails"
)

func KillmailsList() ([]kmm.Killmail, error) {
	results, err := kmm.GetList()

	if err != nil {
		return nil, err
	}

	return results, nil
}
