package main

import (
	"fmt"
	"github.com/dalton-hill/osrs-xp-calc/actions"
	"github.com/dalton-hill/osrs-xp-calc/actions/ore"
	"github.com/dalton-hill/osrs-xp-calc/items"
)

func main() {
	// sort all ore bars by xp per coal used
	barActions := ore.GenerateActionSlice()
	coalActions := actions.FilterActions(barActions, actions.RequiresItemFilter(items.COAL))
	coalActions.SortByXpPer(items.COAL)
	fmt.Println(coalActions)
	// smelt in sorted order until resources depleted
	// smelt all gold/silver based on highest xp rates (add glove toggle later)

}