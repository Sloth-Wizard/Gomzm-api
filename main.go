package main

import (
	"fmt"
	"net/http"

	API "gomzm-api/api/killmails"
)

/*
Call to API => list of killmails
*/
func getKmList(w http.ResponseWriter, r *http.Request) {
	killmails, err := API.KillmailsList()
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(killmails)
}

/*
Manage all the API routing
*/
func main() {
	_, err := API.KillmailsList()
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		fmt.Printf("[%s] Starting API ...\n", time.Now().Format("2006-01-02 15:04:05"))
		mux := http.NewServeMux()
		mux.HandleFunc("/api/killmails/", getKmList)

		s := &http.Server{
			Addr:    ":1337",
			Handler: mux,
		}
		s.ListenAndServe()
	*/
}
