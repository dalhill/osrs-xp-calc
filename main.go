package main

import (
	"fmt"
	"github.com/dalton-hill/osrs-xp-calc/actions"
	"github.com/dalton-hill/osrs-xp-calc/items"
	"github.com/dalton-hill/osrs-xp-calc/modifications"
	"github.com/dalton-hill/osrs-xp-calc/skills"
)


/*
	todo:
		- allow user to not take an aciton, for example don't make iron bars because user only wants to use iron ore for steel bars
		- allow user to assign priority, ex: {action: MakeIron, priority: 8}
		- only thing to add after that would be applying boons (ex: gold gauntlets)

		algorithm as follows
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


todo: food for thought...
	1. either take maxActions OR actions require for next level
	2. if leveled up in step 1 then recalculate everything as we may have access to new methods
*/


// todo: for simplicity sake, potion input/output will be 1 count per one dose. Ex: making pray potion produces 3 prayer potions
//  making super combat requires 4 super attack, 4 super strength, 4 super defence
//  when displaying total potions made, just divide count by 4 and show them 4 dose potions...

func main() {
	itemMap := items.LoadItemsFromJSON("items/items.json")
	actionSlice := actions.LoadActionsFromJSON("actions/actions.json")
	actionSlice.SortByXpPer(items.Coal, skills.Smithing)
	modificationSlice := []modifications.Modification{modifications.BlastFurnace, modifications.GoldGauntlets} // todo: load & filter to user selected

	// apply modifications
	for i, a := range actionSlice {
		actionSlice[i] = modifications.ApplyModifications(a, modificationSlice)
	}

	// take action
	for i := range actionSlice {
		actions.TakeMaxAction(&actionSlice[i], itemMap)
	}

	// display to user
	// fmt.Println("items remaining: ", itemMap)
	fmt.Println("experience outputs: ")
	skillExperience := actionSlice.GetTotalXP()
	for k, v := range skillExperience {
		fmt.Printf("\t%s: %f\n", k, v)
	}
}

