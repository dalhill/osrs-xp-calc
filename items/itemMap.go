package items

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type ItemMap map[string]int

// LoadFromJSON loads ItemMap from the specified path.
func LoadItemsFromJSON(filename string) ItemMap {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var im ItemMap
		if err := json.Unmarshal(bs, &im); err != nil {
		log.Fatal(err)
	}
	return im
}