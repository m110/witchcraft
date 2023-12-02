package component

import (
	"github.com/yohamta/donburi"

	"github.com/m110/witchcraft/engine"
)

type ManaData struct {
	Mana    int
	MaxMana int

	ManaRegenTimer *engine.Timer
	ManaRegen      int
}

func (d *ManaData) UseMana(amount int) bool {
	if d.Mana-amount < 0 {
		return false
	}

	d.Mana -= amount
	return true
}

func (d *ManaData) AddMana(amount int) {
	d.Mana += amount
	if d.Mana > d.MaxMana {
		d.Mana = d.MaxMana
	}
}

var Mana = donburi.NewComponentType[ManaData]()
