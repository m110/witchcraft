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
	for _, a := range d.Auras {
		if a.Equals(aura) {
			return
		}
	}

	d.Auras = append(d.Auras, aura)
}

var AuraHolder = donburi.NewComponentType[AuraHolderData]()

type Aura struct {
	Effect spell.AuraEffect
	Source donburi.Entity

	Timer     *engine.Timer
	TickTimer *engine.Timer
}

func NewAura(source *donburi.Entry, effect spell.AuraEffect) Aura {
	a := Aura{
		Effect: effect,
		Source: source.Entity(),
	}

	if effect.Duration != 0 {
		a.Timer = engine.NewTimer(effect.Duration)
	}

	if effect.TickTime != 0 {
		a.TickTimer = engine.NewTimer(effect.TickTime)
	}

	return a
}

func (a *Aura) Equals(other Aura) bool {
	return a.Effect.ID == other.Effect.ID && a.Source == other.Source
}

func (a *Aura) EqualsTo(source *donburi.Entry, effect spell.AuraEffect) bool {
	return a.Effect.ID == effect.ID && a.Source == source.Entity()
}
