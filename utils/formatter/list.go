package formatter

import (
	"encoding/json"
	"fmt"
)

func loop(slice []interface{}) {
	for _, elem := range slice {
		switch elemTyped := elem.(type) {
		case int:
			fmt.Println("int:", elemTyped)
		case string:
			fmt.Println("string:", elemTyped)
		case []string:
			fmt.Println("[]string:", elemTyped)
		case interface{}:
			fmt.Println("map:", elemTyped)
		}
	}
}

/*
Setup the right format for the list endpoint, also transforms the data to bytes
*/
func KmToList(data interface{}) ([]byte, error) {
	kml, err := json.Marshal(data)
	if err != nil {
		return []byte("null"), err
	}

	kmlJson := string(kml[:])
	var unMarchalledJSON []interface{}
	json.Unmarshal([]byte(kmlJson), &unMarchalledJSON)
	loop(unMarchalledJSON)

	return kml, nil
}
