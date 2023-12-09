package archetype

import (
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
				Body:       component.AllBodyParts.Bodies[0],
				Hair:       &component.AllBodyParts.Hairs[44],
				FacialHair: &component.AllBodyParts.FacialHairs[18],
				Equipment: component.Equipment{
					Head:     &component.AllBodyParts.HeadArmors[33],
					Chest:    &component.AllBodyParts.ChestArmors[5],
					Legs:     &component.AllBodyParts.LegsArmors[0],
					Feet:     &component.AllBodyParts.FeetArmors[0],
					MainHand: &component.AllBodyParts.MainHandWeapons[12],
					OffHand:  nil,
				},
			},
			Spells: []component.Spell{
				component.NewSpell(spell.FireBall),
				component.NewSpell(spell.LightningBolt),
				component.NewSpell(spell.ManaSurge),
			},
		},
		{
			Name: "Wizard",
			Character: component.CharacterData{
				Body:       component.AllBodyParts.Bodies[2],
				Hair:       &component.AllBodyParts.Hairs[21],
				FacialHair: &component.AllBodyParts.FacialHairs[15],
				Equipment: component.Equipment{
					Head:     nil,
					Chest:    &component.AllBodyParts.ChestArmors[1],
					Legs:     nil,
					Feet:     &component.AllBodyParts.FeetArmors[6],
					MainHand: &component.AllBodyParts.MainHandWeapons[0],
					OffHand:  nil,
				},
			},
			Spells: []component.Spell{
				component.NewSpell(spell.Quicksand),
				component.NewSpell(spell.LightningBolt),
				component.NewSpell(spell.Spark),
			},
		},
		{
			Name: "Warlock",
			Character: component.CharacterData{
				Body:       component.AllBodyParts.Bodies[4],
				Hair:       &component.AllBodyParts.Hairs[35],
				FacialHair: &component.AllBodyParts.FacialHairs[13],
				Equipment: component.Equipment{
					Head:     nil,
					Chest:    &component.AllBodyParts.ChestArmors[118],
					Legs:     &component.AllBodyParts.LegsArmors[0],
					Feet:     &component.AllBodyParts.FeetArmors[0],
					MainHand: &component.AllBodyParts.MainHandWeapons[82],
					OffHand:  nil,
				},
			},
			Spells: []component.Spell{
				component.NewSpell(spell.VenomBurst),
				component.NewSpell(spell.LightningBolt),
				component.NewSpell(spell.Spark),
			},
		},
		{
			Name: "Sorcerer",
			Character: component.CharacterData{
				Body:       component.AllBodyParts.Bodies[1],
				Hair:       &component.AllBodyParts.Hairs[11],
				FacialHair: nil,
				Equipment: component.Equipment{
					Head:     &component.AllBodyParts.HeadArmors[35],
					Chest:    &component.AllBodyParts.ChestArmors[84],
					Legs:     nil,
					Feet:     &component.AllBodyParts.FeetArmors[1],
					MainHand: &component.AllBodyParts.MainHandWeapons[36],
					OffHand:  nil,
				},
			},
			Spells: []component.Spell{
				component.NewSpell(spell.ArcaneVolley),
				component.NewSpell(spell.ArcaneBarrage),
				component.NewSpell(spell.Spark),
			},
		},
		{
			Name: "Witch",
			Character: component.CharacterData{
				Body:       component.AllBodyParts.Bodies[3],
				Hair:       &component.AllBodyParts.Hairs[21],
				FacialHair: nil,
				Equipment: component.Equipment{
					Head:     nil,
					Chest:    &component.AllBodyParts.ChestArmors[92],
					Legs:     nil,
					Feet:     nil,
					MainHand: &component.AllBodyParts.MainHandWeapons[0],
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
