/*
Contains all the queries to get killmails from the database
*/
package killmails

import (
	"encoding/json"
	"errors"
	"fmt"
	"gomzm-api/utils/database"
)

const LIMIT = "40" // Must be a string because it's used in queries

type KillmailList struct {
	Killmails []*Killmail `json:"killmails"`
}

type Killmail struct {
	Id                      uint            `json:"id"`
	Killmail_id             uint            `json:"killmail_id"`
	Victim_character_id     uint            `json:"victim_character_id"`
	Final_blow_character_id uint            `json:"final_blow_character_id"`
	Final_blow_faction_id   uint            `json:"final_blow_faction_id"`
	Killmail_details        json.RawMessage `json:"killmail_details"`
}

var (
	db, _ = database.Connect()
)

func Get(_type string) ([]*Killmail, error) {
	switch _type {
	case "list":
		kml := new(KillmailList)
		r, err := kml.GetList()
		if err != nil {
			return nil, err
		}

		return r, nil
	}

	err := errors.New("please specify a type to get (err: 1)")
	return nil, err
}

/*
DB killmail scannable fields
*/
func (k *Killmail) scannableFields() []interface{} {
	return []interface{}{
		&k.Id,
		&k.Killmail_id,
		&k.Victim_character_id,
		&k.Final_blow_character_id,
		&k.Final_blow_faction_id,
		&k.Killmail_details}
}

/*
Query the DB and get a list of killmails
*/
func (kml *KillmailList) GetList() ([]*Killmail, error) {
	k := new(Killmail)

	// Prepare and execute the statement
	fmt.Printf("Getting KM from DB ...\n")

	q := "SELECT * FROM killmails LIMIT ?"
	rows, err := db.Query(q, LIMIT)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Scan results
	for rows.Next() {
		if err := rows.Scan(k.scannableFields()...); err != nil {
			return nil, err
		}

		kml.Killmails = append(kml.Killmails, k)
	}

	return kml.Killmails, nil
}
