package actions

import (
	"fmt"
	"github.com/dalton-hill/osrs-xp-calc/items"
	"sort"
)

type Action struct {
	Name string
	XpReward float64
	RequiredResources []items.ItemCount
}

type ActionSlice []Action

type ActionFilter func(Action) bool

type ActionSliceComparison func(ActionSlice) func(i, j int) bool

func RequiresItemFilter(itemName string) func(Action) bool {
	return func(a Action) bool {
		for _, ic := range a.RequiredResources {
			if ic.ItemName == itemName {
				return true
			}
		}
		return false
	}
}

func FilterActions(as ActionSlice, f ActionFilter) ActionSlice {
	var cs ActionSlice
	for _, a := range as {
		if f(a) {
			cs = append(cs, a)
		}
	}
	return cs
}

func (as ActionSlice) SortByXpPer(reqName string) {
	sort.Slice(as[:], func(i, j int) bool {
		return as[i].XpPer(reqName) > as[j].XpPer(reqName)
	})
}

func (a Action) XpPer(reqName string) float64 {
	for _, item := range a.RequiredResources {
		if item.ItemName == reqName {
			fmt.Println(a.Name, a.XpReward / float64(item.Count))
			return a.XpReward / float64(item.Count)
		}
	}
	return 0
}