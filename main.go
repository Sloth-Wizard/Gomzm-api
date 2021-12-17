package main

import (
	"fmt"
	"gomzm-api/api/killmails"
)

func main() {
	killmails, _ := killmails.KillmailsList()
	fmt.Println(killmails)
}
