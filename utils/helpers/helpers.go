package helpers

import "encoding/json"

/*
Gnnnnnnn krmzm, gooby pls
*/
func Goobynator(boring_str string) string {
	return boring_str
}

/*
Pretty print stuff
*/
func PPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
