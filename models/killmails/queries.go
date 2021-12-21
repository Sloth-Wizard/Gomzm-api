/*
Contains all the queries to get killmails from the database
*/
package killmails

import (
	"encoding/json"
	"errors"
	"fmt"
	"gomzm-api/utils/database"
	"time"
)

type Killmail struct {
	Id                      uint            `json:"id"`
	Killmail_id             uint            `json:"killmail_id"`
	Victim_character_id     uint            `json:"victim_character_id"`
	Final_blow_character_id uint            `json:"final_blow_character_id"`
	Final_blow_faction_id   uint            `json:"final_blow_faction_id"`
	Killmail_details        json.RawMessage `json:"killmail_details"`
}

type KillmailList struct {
	Killmails []*Killmail `json:"killmails"`
}

var (
	db, _ = database.Connect()
)

func Get(_type string) ([]*Killmail, error) {
	switch _type {
	case "list":
		kml := new(KillmailList)
		r, err := kml.getList()
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
func (kml *KillmailList) getList() ([]*Killmail, error) {
	// Prepare and execute the statement
	fmt.Printf("[%s] Getting KM from DB ...\n", time.Now().Format("2006-01-02 15:04:05"))

	q := "SELECT * FROM killmails LIMIT ?"
	rows, err := db.Query(q, 40)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Scan results
	for rows.Next() {
		k := new(Killmail)
		if err := rows.Scan(k.scannableFields()...); err != nil {
			return nil, err
		}

		kml.Killmails = append(kml.Killmails, k)
	}

	return kml.Killmails, nil
}
