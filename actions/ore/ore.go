package ore

import (
	"github.com/dalton-hill/osrs-xp-calc/actions"
	"encoding/json"
	"io/ioutil"
	"log"
)

const MakeBronze = "MakeBronze"
const MakeIron = "MakeIron"
const MakeSilver = "MakeSilver"
const MakeSteel = "MakeSteel"
const MakeGold = "MakeGold"
const MakeMithril = "MakeMithril"
const MakeAdamant = "MakeAdamant"
const MakeRune = "MakeRune"

// LoadFromJSON loads actions.ActionSlice from the specified path.
func LoadFromJSON(filename string) actions.ActionSlice {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var as actions.ActionSlice
	if err := json.Unmarshal(bs, &as); err != nil {
		log.Fatal(err)
	}
	return as
}