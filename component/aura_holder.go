package component

import (
	"github.com/m110/witchcraft/engine"
	"github.com/m110/witchcraft/spell"
	"github.com/yohamta/donburi"
)

type AuraHolderData struct {
	Auras []Aura
}

func (d *AuraHolderData) ApplyAura(aura Aura) {
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
