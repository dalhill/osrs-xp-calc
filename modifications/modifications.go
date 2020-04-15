package modifications

import (
	"fmt"
	"github.com/dalton-hill/osrs-xp-calc/actions"
	"github.com/dalton-hill/osrs-xp-calc/items"
)

type UserSelection struct {
	ModificationName string
	ShouldApply      bool
}

// not sure exactly how i want to implement this yet
type Modification struct {
	Name        string
	ActionNames []string
	ShouldApply bool
	Modify      func(actions.Action) actions.Action
}

// Modifier will be something such as GoldGauntlets which would return an
// updated MAKE_GOLD action which has its experience reward doubled.
// type Modifier interface {
// 	Modify func
// }

var GoldGauntlets = Modification{
	Name:        "GoldGauntlets",
	ActionNames: []string{actions.MAKE_GOLD},
	ShouldApply: false,
	Modify: func(a actions.Action) actions.Action {
		a.XpReward = 55
		return a
	},
}

var BlastFurnace = Modification{
	Name:        "BlastFurnace",
	ActionNames: []string{actions.MAKE_STEEL, actions.MAKE_MITHRIL, actions.MAKE_ADAMANT, actions.MAKE_RUNE},
	ShouldApply: false,
	Modify: func(a actions.Action) actions.Action {
		for i := range a.RequiredResources {
			r := &a.RequiredResources[i]
			if r.Name == items.COAL {
				r.Count = r.Count / 2
			}
		}
		return a
	},
}

func (m Modification) CanModify(a actions.Action) bool {
	for _, aName := range m.ActionNames {
		if aName == a.Name {
			fmt.Printf("%s can modify %s\n", m.Name, a.Name)
			return true
		}
	}
	return false
}
