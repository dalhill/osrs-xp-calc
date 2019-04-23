package actions

import (
	"github.com/dalton-hill/osrs-xp-calc/items"
)

type Action struct {
	Name string
	XpReward float64
	RequiredResources []items.ItemCount
}

type ActionSlice []Action

type ActionFilter func(Action) bool

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

func FilterActions(as ActionSlice, f ActionFilter) []Action {
	var cs []Action
	for _, a := range as {
		if f(a) {
			cs = append(cs, a)
		}
	}
	return cs
}