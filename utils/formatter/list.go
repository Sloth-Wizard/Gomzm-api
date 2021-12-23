package formatter

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ApiKmList struct {
	Killmail_id                       uint    `json:"killmail_id"`
	Time                              string  `json:"time"`
	Ship_image                        string  `json:"ship_image"`
	Victim_name                       string  `json:"victim_name"`
	Victim_gooby_name                 string  `json:"victim_gooby_name"`
	Victim_corporation                string  `json:"victim_corporation"`
	Victim_gooby_corporation          string  `json:"victim_gooby_corporation"`
	Victim_corporation_image          string  `json:"victim_corporation_image"`
	Victim_alliance                   string  `json:"victim_alliance"`
	Victim_gooby_alliance             string  `json:"victim_gooby_alliance"`
	Victim_alliance_image             string  `json:"victim_alliance_image"`
	Final_blow_name                   string  `json:"final_blow_name"`
	Final_blow_gooby_name             string  `json:"final_blow_gooby_name"`
	Final_blow_faction                string  `json:"final_blow_faction"`
	Final_blow_faction_name           string  `json:"final_blow_faction_name"`
	Final_blow_faction_gooby_name     string  `json:"final_blow_faction_gooby_name"`
	Final_blow_faction_image          string  `json:"final_blow_faction_image"`
	Final_blow_corporation            uint    `json:"final_blow_corporation"`
	Final_blow_corporation_name       string  `json:"final_blow_corporation_name"`
	Final_blow_corporation_gooby_name string  `json:"final_blow_corporation_gooby_name"`
	Final_blow_corporation_image      string  `json:"final_blow_corporation_image"`
	Final_blow_alliance               uint    `json:"final_blow_alliance"`
	Final_blow_alliance_name          string  `json:"final_blow_alliance_name"`
	Final_blow_alliance_gooby_name    string  `json:"final_blow_alliance_gooby_name"`
	Final_blow_alliance_image         string  `json:"final_blow_alliance_image"`
	Involved                          uint    `json:"involved"`
	Total_price                       float64 `json:"total_price"`
}

func (akl *ApiKmList) structBuilder(identifier string, data interface{}) {
	err := "Error adding data to struct, wrong type given"
	switch identifier {
	case "Killmail_id":
		if v, ok := data.(float64); ok {
			akl.Killmail_id = uint(v)
		} else {
			fmt.Println(err)
		}
	case "Ship_image":
		if v, ok := data.(string); ok {
			akl.Ship_image = string(v)
		} else {
			fmt.Println(err)
		}
	}

	fmt.Println(akl)
}

/*
Setup the right format for the list endpoint, also transforms the data to bytes
*/
func KmToList(data interface{}) ([]byte, error) {
	kml, err := json.Marshal(data)
	if err != nil {
		return []byte(nil), err
	}

	var akl []interface{}
	if err := json.Unmarshal(kml, &akl); err != nil {
		return []byte(nil), err
	}

	var result []interface{}
	for _, km := range akl {
		fmt.Printf("\n---------------------------------------------------------\n\n")

		if killmail, ok := km.(map[string]interface{}); ok {
			item := new(ApiKmList)
			for k, v := range killmail {
				fmt.Printf(" [========>] %s = %s \n", k, v)
				item.structBuilder(strings.Title(k), v)
			}
			result = append(result, item)
		} else {
			fmt.Printf("record not a map[string]interface{}: %v\n", km)
		}

		fmt.Printf("\n---------------------------------------------------------\n\n")
	}

	fmt.Println(&result)
	/*
		kmlJson := string(kml[:])
		var unMarchalledJSON []interface{}
		json.Unmarshal([]byte(kmlJson), &unMarchalledJSON)
	*/

	return kml, nil
}
