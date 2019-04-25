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

func main() {
	itemStore := items.ItemSlice{
		{
			Name: items.IRON_ORE,
			Count: 12675,
		},
		{
			Name: items.SILVER_ORE,
			Count: 566,
		},
		{
			Name: items.COAL,
			Count: 12014,
		},
		{
			Name: items.GOLD_ORE,
			Count: 8453,
		},
		{
			Name: items.MITHRIL_ORE,
			Count: 5860,
		},
		{
			Name: items.ADAMANT_ORE,
			Count: 5837,
		},
		{
			Name: items.RUNE_ORE,
			Count: 423,
		},
	}
	actionStore := ore.GenerateActionSlice()
	actionStore.SortByXpPer(items.COAL)

	for i := range actionStore {
		actions.TakeMaxAction(&actionStore[i], itemStore)
	}
	fmt.Println(itemStore)
	fmt.Println(actionStore)
}