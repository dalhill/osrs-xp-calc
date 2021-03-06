package modifications

import (
	"fmt"
	"github.com/dalton-hill/osrs-xp-calc/actions"
	"github.com/dalton-hill/osrs-xp-calc/items"
	"github.com/dalton-hill/osrs-xp-calc/skills"
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
// updated MakeGold action which has its experience reward doubled.
// type Modifier interface {
// 	Modify func
// }

var GoldGauntlets = Modification{
	Name:        "GoldGauntlets",
	ActionNames: []string{actions.MakeGold},
	ShouldApply: false,
	Modify: func(a actions.Action) actions.Action {
		a.ExperienceOutputs[skills.Smithing] = 55
		return a
	},
}

var BlastFurnace = Modification{
	Name:        "BlastFurnace",
	ActionNames: []string{actions.MakeSteel, actions.MakeMithril, actions.MakeAdamant, actions.MakeRune},
	ShouldApply: false,
	Modify: func(a actions.Action) actions.Action {
		if count, ok := a.ItemInputs[items.Coal]; ok {  // todo: can we get pointer here instead of double getting?
			a.ItemInputs[items.Coal] = count / 2
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

func ApplyModifications(a actions.Action, ms []Modification) actions.Action {
	for _, m := range ms {
		if m.CanModify(a) {
			a = m.Modify(a)
		}
	}
	return a
}
