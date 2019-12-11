package main

import (
	"fmt"

	"github.com/dalton-hill/osrs-xp-calc/actions"
	"github.com/dalton-hill/osrs-xp-calc/actions/ore"
	"github.com/dalton-hill/osrs-xp-calc/items"
)

/*
   iron: 12675
   adamant: 5837
   rune: 423
   silver: 566
   coal: 12014
   mithril: 5860
   gold: 8453
*/

/*
	todo:
		- load store from json config instead of having it exist in code
		- allow user to not take an aciton, for example don't make iron bars because user only wants to use iron ore for steel bars

		- allow user to assign priority, ex: {action: makeIron, priority: 8}
		- only thing to add after that would be applying boons (ex: gold gauntlets)
*/

func main() {
	itemStore := items.LoadFromJSON("bank.json")
	actionStore := ore.LoadFromJSON("actions/ore/ore.json")
	actionStore.SortByXpPer(items.COAL)

	for i := range actionStore {
		actions.TakeMaxAction(&actionStore[i], itemStore)
	}
	fmt.Println(itemStore)
	fmt.Printf("TotalXP: %f\n", actionStore.GetTotalXP())
}
