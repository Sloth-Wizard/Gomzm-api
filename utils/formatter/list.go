package formatter

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type ApiKmList struct {
	Killmail_id                       uint    `json:"killmail_id"`
	Time                              string  `json:"time"`
	Ship_image                        string  `json:"ship_image"`
	Victim_name                       string  `json:"victim_name"`
	Victim_gooby_name                 string  `json:"victim_gooby_name"`
	Victim_corporation                uint    `json:"victim_corporation"`
	Victim_gooby_corporation          string  `json:"victim_gooby_corporation"`
	Victim_corporation_image          string  `json:"victim_corporation_image"`
	Victim_alliance                   string  `json:"victim_alliance"`
	Victim_gooby_alliance             string  `json:"victim_gooby_alliance"`
	Victim_alliance_image             string  `json:"victim_alliance_image"`
	Final_blow_name                   string  `json:"final_blow_name"`
	Final_blow_gooby_name             string  `json:"final_blow_gooby_name"`
	Final_blow_faction                uint    `json:"final_blow_faction"`
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
	Involved                          int     `json:"involved"`
	Total_price                       float64 `json:"total_price"`
}

func (akl *ApiKmList) getFinalBlow(data interface{}) error {
	attackers, err := json.Marshal(data)
	if err != nil { // No attackers ?
		return err
	}

	var r []interface{}
	if err := json.Unmarshal(attackers, &r); err != nil {
		return err
	}

	for _, a_v := range r {
		trigger := false

		if attacker, ok := a_v.(map[string]interface{}); ok {
			for k, v := range attacker {
				if k == "final_blow" && v == true {
					trigger = true
				}
			}

			if trigger {
				for fb_k, fb_v := range attacker {
					fmt.Printf("	\033[32m[========>]\033[0m \033[36m%s\033[0m = %s \n\n", fb_k, fb_v)
					switch fb_k { // TODO: add faction exception if happens
					case "name":
						akl.Final_blow_name = fb_v.(string)
					case "gooby_name":
						akl.Final_blow_gooby_name = fb_v.(string)
					case "faction_name":
						akl.Final_blow_faction_name = fb_v.(string)
					case "faction_gooby_name":
						akl.Final_blow_faction_gooby_name = fb_v.(string)
					case "faction_image":
						akl.Final_blow_faction_image = fb_v.(string)
					case "corporation_id":
						if reflect.TypeOf(fb_v).String() == "string" {
							i, err := strconv.Atoi(fb_v.(string))
							if err != nil {
								return err
							}
							akl.Final_blow_faction, akl.Final_blow_corporation = uint(i), uint(i)
						} else {
							akl.Final_blow_corporation = uint(fb_v.(float64))
						}
					case "corporation_name":
						akl.Final_blow_corporation_name = fb_v.(string)
					case "corporation_gooby_name":
						akl.Final_blow_corporation_gooby_name = fb_v.(string)
					case "corporation_image":
						akl.Final_blow_corporation_image = fb_v.(string)
					case "alliance_id":
						akl.Final_blow_alliance = uint(fb_v.(float64))
					case "alliance_name":
						akl.Final_blow_alliance_name = fb_v.(string)
					case "alliance_gooby_name":
						akl.Final_blow_alliance_gooby_name = fb_v.(string)
					case "alliance_image":
						akl.Final_blow_alliance_image = fb_v.(string)
					}
				}
			}
		} else { // No attacker ?
			return errors.New("no attackers ? skip this kill it isn't valid")
		}
	}

	akl.Involved = len(r)

	return nil
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
		out:
			for k, v := range killmail {
				if k == "killmail_details" {
					if killmail_details, ok := v.(map[string]interface{}); ok { // Split details into victim and attackers
						for kd_k, kd_v := range killmail_details {
							switch kd_k {
							case "attackers": // Get the final blow attacker
								if err := item.getFinalBlow(kd_v); err != nil { // Inject final blow data into the struct
									break out // No attackers ? skip this killmail then (should never happen)
								} else {
									fmt.Println("final blow data")
									fmt.Println(item)
								}
							case "victim": // Split all victim entries
								fmt.Println("victim")
							}
						}
					} else {
						fmt.Printf("record not a map[string]interface{}: %v\n", v)
					}
				} else {
					fmt.Printf("\033[32m[========>]\033[0m \033[36m%s\033[0m = %s \n\n", k, v)
				}
			}
			result = append(result, item)
		} else {
			fmt.Printf("\033[33mrecord not a map[string]interface{}: %v\n\033[0m", km)
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
