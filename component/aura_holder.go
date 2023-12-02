package component

import (
	"github.com/yohamta/donburi"

	"github.com/m110/witchcraft/engine"
	"github.com/m110/witchcraft/spell"
)

type AuraHolderData struct {
	Auras []Aura

	UI *donburi.Entry
}

func (d *AuraHolderData) ApplyAura(aura Aura) {
	// TODO Check stacking / uniqueness
	d.Auras = append(d.Auras, aura)
}

var AuraHolder = donburi.NewComponentType[AuraHolderData]()

type Aura struct {
	Template spell.AuraEffect

	Timer     *engine.Timer
	TickTimer *engine.Timer
}

func NewAura(template spell.AuraEffect) Aura {
	return Aura{
		Template:  template,
		Timer:     engine.NewTimer(template.Duration),
		TickTimer: engine.NewTimer(template.TickTime),
	}
}
