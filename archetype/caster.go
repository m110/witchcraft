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
				Body:       assets.RandomFrom(assets.Bodies),
				Hair:       assets.RandomFromOrEmpty(assets.Hairs),
				FacialHair: assets.RandomFromOrEmpty(assets.FacialHairs),
				Equipment: component.Equipment{
					Head:     assets.RandomFromOrEmpty(assets.HeadArmors),
					Chest:    assets.RandomFromOrEmpty(assets.ChestArmors),
					Legs:     assets.RandomFromOrEmpty(assets.LegsArmors),
					Feet:     assets.RandomFromOrEmpty(assets.FeetArmors),
					MainHand: assets.RandomFromOrEmpty(assets.MainHandWeapons),
					OffHand:  assets.RandomFromOrEmpty(assets.OffHandWeapons),
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
				Body:       assets.RandomFrom(assets.Bodies),
				Hair:       assets.RandomFromOrEmpty(assets.Hairs),
				FacialHair: assets.RandomFromOrEmpty(assets.FacialHairs),
				Equipment: component.Equipment{
					Head:     assets.RandomFromOrEmpty(assets.HeadArmors),
					Chest:    assets.RandomFromOrEmpty(assets.ChestArmors),
					Legs:     assets.RandomFromOrEmpty(assets.LegsArmors),
					Feet:     assets.RandomFromOrEmpty(assets.FeetArmors),
					MainHand: assets.RandomFromOrEmpty(assets.MainHandWeapons),
					OffHand:  assets.RandomFromOrEmpty(assets.OffHandWeapons),
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
				Body:       assets.RandomFrom(assets.Bodies),
				Hair:       assets.RandomFromOrEmpty(assets.Hairs),
				FacialHair: assets.RandomFromOrEmpty(assets.FacialHairs),
				Equipment: component.Equipment{
					Head:     assets.RandomFromOrEmpty(assets.HeadArmors),
					Chest:    assets.RandomFromOrEmpty(assets.ChestArmors),
					Legs:     assets.RandomFromOrEmpty(assets.LegsArmors),
					Feet:     assets.RandomFromOrEmpty(assets.FeetArmors),
					MainHand: assets.RandomFromOrEmpty(assets.MainHandWeapons),
					OffHand:  assets.RandomFromOrEmpty(assets.OffHandWeapons),
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
				Body:       assets.RandomFrom(assets.Bodies),
				Hair:       assets.RandomFromOrEmpty(assets.Hairs),
				FacialHair: assets.RandomFromOrEmpty(assets.FacialHairs),
				Equipment: component.Equipment{
					Head:     assets.RandomFromOrEmpty(assets.HeadArmors),
					Chest:    assets.RandomFromOrEmpty(assets.ChestArmors),
					Legs:     assets.RandomFromOrEmpty(assets.LegsArmors),
					Feet:     assets.RandomFromOrEmpty(assets.FeetArmors),
					MainHand: assets.RandomFromOrEmpty(assets.MainHandWeapons),
					OffHand:  assets.RandomFromOrEmpty(assets.OffHandWeapons),
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
				Body:       assets.RandomFrom(assets.Bodies),
				Hair:       assets.RandomFromOrEmpty(assets.Hairs),
				FacialHair: assets.RandomFromOrEmpty(assets.FacialHairs),
				Equipment: component.Equipment{
					Head:     assets.RandomFromOrEmpty(assets.HeadArmors),
					Chest:    assets.RandomFromOrEmpty(assets.ChestArmors),
					Legs:     assets.RandomFromOrEmpty(assets.LegsArmors),
					Feet:     assets.RandomFromOrEmpty(assets.FeetArmors),
					MainHand: assets.RandomFromOrEmpty(assets.MainHandWeapons),
					OffHand:  assets.RandomFromOrEmpty(assets.OffHandWeapons),
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
