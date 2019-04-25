package ore

import (
	"github.com/dalton-hill/osrs-xp-calc/actions"
	. "github.com/dalton-hill/osrs-xp-calc/items"
)

const MakeBronze = "MakeBronze"
const MakeIron = "MakeIron"
const MakeSilver = "MakeSilver"
const MakeSteel = "MakeSteel"
const MakeGold = "MakeGold"
const MakeMithril = "MakeMithril"
const MakeAdamant = "MakeAdamant"
const MakeRune = "MakeRune"

func GenerateActionSlice() actions.ActionSlice {
	return actions.ActionSlice{
		{
			Name:     MakeBronze,
			XpReward: 6.5,
			RequiredResources: []Item{
				{
					Name: COPPER_ORE,
					Count: 1,
				},
				{
					Name: TIN_ORE,
					Count: 1,
				},
			},
		},
		{
			Name:     MakeIron,
			XpReward: 12.5,
			RequiredResources: []Item{
				{
					Name: IRON_ORE,
					Count: 1,
				},
			},
		},
		{
			Name:     MakeSteel,
			XpReward: 17.5,
			RequiredResources: []Item{
				{
					Name: IRON_ORE,
					Count: 1,
				},
				{
					Name: COAL,
					Count: 2,
				},
			},
		},
		{
			Name:     MakeSilver,
			XpReward: 13.7,
			RequiredResources: []Item{
				{
					Name: SILVER_ORE,
					Count: 1,
				},
			},
		},
		{
			Name:     MakeMithril,
			XpReward: 30,
			RequiredResources: []Item{
				{
					Name: MITHRIL_ORE,
					Count: 1,
				},
				{
					Name: COAL,
					Count: 4,
				},
			},
		},
		{
			Name:     MakeGold,
			XpReward: 22.5,
			RequiredResources: []Item{
				{
					Name: GOLD_ORE,
					Count: 1,
				},
			},
		},
		{
			Name:     MakeAdamant,
			XpReward: 37.5,
			RequiredResources: []Item{
				{
					Name: ADAMANT_ORE,
					Count: 1,
				},
				{
					Name: COAL,
					Count: 6,
				},
			},
		},
		{
			Name:     MakeRune,
			XpReward: 75,
			RequiredResources: []Item{
				{
					Name: RUNE_ORE,
					Count: 1,
				},
				{
					Name: COAL,
					Count: 8,
				},
			},
		},
	}
}