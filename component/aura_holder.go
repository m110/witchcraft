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
	Source   donburi.Entity

	Timer     *engine.Timer
	TickTimer *engine.Timer
}

func NewAura(source *donburi.Entry, template spell.AuraEffect) Aura {
	a := Aura{
		Template: template,
		Source:   source.Entity(),
	}

	if template.Duration != 0 {
		a.Timer = engine.NewTimer(template.Duration)
	}

	if template.TickTime != 0 {
		a.TickTimer = engine.NewTimer(template.TickTime)
	}

	return a
}

func (a *Aura) Equals(other Aura) bool {
	return a.Template.ID == other.Template.ID && a.Source == other.Source
}
