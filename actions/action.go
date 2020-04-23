/*
things to consider
	will the user sell the bars or smith them?

ultimately should make an MVP that makes a few assumptions
*/
package actions

import (
	"encoding/json"
	"github.com/dalton-hill/osrs-xp-calc/items"
	"io/ioutil"
	"log"
	"math"
	"sort"
)

type Action struct {
	Name              string
	Count             int
	ItemInputs        items.ItemMap
	ItemOutputs       items.ItemMap
	ExperienceOutputs map[string]float64
}

type ActionSlice []Action

type ActionFilter func(Action) bool

type ActionSliceComparison func(ActionSlice) func(i, j int) bool

func (as ActionSlice) SortByXpPer(reqName string, skillName string) {
	sort.Slice(as[:], func(i, j int) bool {
		return as[i].XpPer(reqName, skillName) > as[j].XpPer(reqName, skillName)
	})
}

func (a Action) XpPer(reqName string, skillName string) float64 {
	if count, ok := a.ItemInputs[reqName]; ok {
		if xp, ok := a.ExperienceOutputs[skillName]; ok {
			return xp / float64(count)
		}
	}
	return 0
}

// Returns the maximum number of times the action can be performed given the itemMap
func MaxActionCount(a Action, itemMap items.ItemMap) int {
	validCount := false
	maxCount := math.MaxInt32
	for req := range a.ItemInputs {
		if numItems, ok := itemMap[req]; ok {
			validCount = true
			count := numItems / a.ItemInputs[req]
			if count < maxCount {
				maxCount = count
			}
		}
	}
	if validCount {
		return maxCount
	}
	return 0
}

// GetTotalXP sums up the total XP gain that is represented by the ActionSlice
func (as ActionSlice) GetTotalXP() map[string]float64 {
	skillExperience := make(map[string]float64)
	for _, a := range as {
		for k, v := range a.ExperienceOutputs {
			e := float64(a.Count) * v
			if _, ok := skillExperience[k]; ok {
				skillExperience[k] += e
			} else {
				skillExperience[k] = e
			}

		}
	}
	return skillExperience
}

// LoadFromJSON loads actions.ActionSlice from the specified path.
func LoadActionsFromJSON(filename string) ActionSlice {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var as ActionSlice
	if err := json.Unmarshal(bs, &as); err != nil {
		log.Fatal(err)
	}
	return as
}
