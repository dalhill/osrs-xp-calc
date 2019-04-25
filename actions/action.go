package actions

import (
	"fmt"
	"github.com/dalton-hill/osrs-xp-calc/items"
	"sort"
)

type Action struct {
	Name string
	Count int
	XpReward float64
	RequiredResources items.ItemSlice
}

type ActionSlice []Action

type ActionFilter func(Action) bool

type ActionSliceComparison func(ActionSlice) func(i, j int) bool

func RequiresItemFilter(itemName string) func(Action) bool {
	return func(a Action) bool {
		for _, ic := range a.RequiredResources {
			if ic.Name == itemName {
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
		if item.Name == reqName {
			return a.XpReward / float64(item.Count)
		}
	}
	return 0
}

func ResourcesAvailable(a Action, inv []items.Item) bool {
	for _, r := range a.RequiredResources {
		isSatisfied := false
		for _, i := range inv {
			if i.Name == r.Name {
				isSatisfied = i.Count >= r.Count
				break
			}
		}
		if !isSatisfied {
			return false
		}
	}
	return true
}

func TakeMaxAction(a *Action, itemSlice items.ItemSlice) {
	// reqIndex, itemSliceIndex
	indexMap := items.IndexMap(itemSlice, a.RequiredResources)
	var possibleActions []int
	for k, v := range indexMap {
		current := itemSlice[k].Count / a.RequiredResources[v].Count
		possibleActions = append(possibleActions, current)
	}
	var maxAction int
	if len(possibleActions) > 0 {
		maxAction = possibleActions[0]
		for _, v := range possibleActions {
			if v < maxAction {
				maxAction = v
			}
		}
	}
	for k, v := range indexMap {
		itemSlice[k].Count -= a.RequiredResources[v].Count * maxAction
	}
	a.Count += maxAction
	fmt.Println(a.Name, a.Count)
}

