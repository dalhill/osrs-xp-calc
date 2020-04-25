package actions

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Smithing
const MakeSteel = "MakeSteel"
const MakeMithril = "MakeMithril"
const MakeAdamant = "MakeAdamant"
const MakeRune = "MakeRune"
const MakeGold = "MakeGold"

func LoadNamesFromJSON(filename string) []string {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var names []string
	if err := json.Unmarshal(bs, &names); err != nil {
		log.Fatal(err)
	}
	return names
}
