package archetype

import (
	"github.com/m110/witchcraft/assets"
	"github.com/m110/witchcraft/component"
	"github.com/m110/witchcraft/spell"
)

type Class struct {
	Name      string
	Character component.CharacterData
	Spells    []component.Spell
}

func LoadClasses() []Class {
	return []Class{
		{
			Name: "Mage",
			Character: component.CharacterData{
				Body:       assets.Bodies[0],
				Hair:       &assets.Hairs[44],
				FacialHair: &assets.FacialHairs[18],
				Equipment: component.Equipment{
					Head:     &assets.HeadArmors[33],
					Chest:    &assets.ChestArmors[5],
					Legs:     &assets.LegsArmors[0],
					Feet:     &assets.FeetArmors[0],
					MainHand: &assets.MainHandWeapons[12],
					OffHand:  nil,
				},
			},
			Spells: []component.Spell{
				component.NewSpell(spell.FireBall),
				component.NewSpell(spell.LightningBolt),
				component.NewSpell(spell.Spark),
			},
		},
		{
			Name: "Wizard",
			Character: component.CharacterData{
				Body:       assets.Bodies[2],
				Hair:       &assets.Hairs[21],
				FacialHair: &assets.FacialHairs[15],
				Equipment: component.Equipment{
					Head:     nil,
					Chest:    &assets.ChestArmors[1],
					Legs:     nil,
					Feet:     &assets.FeetArmors[6],
					MainHand: &assets.MainHandWeapons[0],
					OffHand:  nil,
				},
			},
			Spells: []component.Spell{
				component.NewSpell(spell.FireBall),
				component.NewSpell(spell.LightningBolt),
				component.NewSpell(spell.Spark),
			},
		},
		{
			Name: "Warlock",
			Character: component.CharacterData{
				Body:       assets.Bodies[4],
				Hair:       &assets.Hairs[35],
				FacialHair: &assets.FacialHairs[13],
				Equipment: component.Equipment{
					Head:     nil,
					Chest:    &assets.ChestArmors[118],
					Legs:     &assets.LegsArmors[0],
					Feet:     &assets.FeetArmors[0],
					MainHand: &assets.MainHandWeapons[82],
					OffHand:  nil,
				},
			},
			Spells: []component.Spell{
				component.NewSpell(spell.FireBall),
				component.NewSpell(spell.LightningBolt),
				component.NewSpell(spell.Spark),
			},
		},
		{
			Name: "Sorcerer",
			Character: component.CharacterData{
				Body:       assets.Bodies[1],
				Hair:       &assets.Hairs[11],
				FacialHair: nil,
				Equipment: component.Equipment{
					Head:     &assets.HeadArmors[35],
					Chest:    &assets.ChestArmors[84],
					Legs:     nil,
					Feet:     &assets.FeetArmors[1],
					MainHand: &assets.MainHandWeapons[36],
					OffHand:  nil,
				},
			},
			Spells: []component.Spell{
				component.NewSpell(spell.FireBall),
				component.NewSpell(spell.LightningBolt),
				component.NewSpell(spell.Spark),
			},
		},
		{
			Name: "Witch",
			Character: component.CharacterData{
				Body:       assets.Bodies[3],
				Hair:       &assets.Hairs[21],
				FacialHair: nil,
				Equipment: component.Equipment{
					Head:     nil,
					Chest:    &assets.ChestArmors[92],
					Legs:     nil,
					Feet:     nil,
					MainHand: &assets.MainHandWeapons[0],
					OffHand:  nil,
				},
			},
			Spells: []component.Spell{
				component.NewSpell(spell.FireBall),
				component.NewSpell(spell.LightningBolt),
				component.NewSpell(spell.Spark),
			},
		},
	}
}
