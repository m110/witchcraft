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

var Mana = donburi.NewComponentType[ManaData]()
