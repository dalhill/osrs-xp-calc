package main

import (
	"fmt"
	"github.com/dalton-hill/osrs-xp-calc/actions"
	"github.com/dalton-hill/osrs-xp-calc/items"
	"github.com/dalton-hill/osrs-xp-calc/modifications"
)

/*
   iron: 12675
   adamant: 5837
   rune: 423
   silver: 566
   Coal: 12014
   mithril: 5860
   gold: 8453
*/

/*
	todo:
		- allow user to not take an aciton, for example don't make iron bars because user only wants to use iron ore for steel bars

		- allow user to assign priority, ex: {action: MakeIron, priority: 8}
		- only thing to add after that would be applying boons (ex: gold gauntlets)

		aglorithm as follows
			1. load config
			2. filter out all inactive modifications
			3. for each action
				1. for each modification
					1. apply if match
			4. filter out all items the user doesn't want selected
			5. sort based on
				- user selected priority
				- OR
				- xp/required resource ratios
			6. take all actions
			7. repeat at step 3 if any actions were taken (items may have been produced)
			8. sum up all xp gained/items produced
*/

func main() {
	itemStore := items.LoadItemsFromJSON("items/items.json")
	actionStore := actions.LoadActionsFromJSON("actions/actions.json")
	actionStore.SortByXpPer(items.Coal)
	mods := []modifications.Modification{modifications.BlastFurnace, modifications.GoldGauntlets} // todo: load & filter to user selected

	for i := range actionStore {
		for _, m := range mods {
			if m.CanModify(actionStore[i]) {
				actionStore[i] = m.Modify(actionStore[i])
			}
		}
		actions.TakeMaxAction(&actionStore[i], itemStore)
	}
	fmt.Println(itemStore)
	fmt.Printf("TotalXP: %f\n", actionStore.GetTotalXP())
}
