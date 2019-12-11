package modifications

import (
	"github.com/dalton-hill/osrs-xp-calc/actions"
	"github.com/dalton-hill/osrs-xp-calc/actions/ore"
)

type UserSelection struct {
	ModificationName string
	ShouldApply      bool
}

// not sure exactly how i want to implement this yet
type Modification struct {
	Name        string
	ActionName  string
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
	ActionName:  ore.MakeGold,
	ShouldApply: false,
	Modify: func(a actions.Action) actions.Action {
		a.XpReward = 55
		return a
	},
}
