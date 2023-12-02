package system

import (
	"fmt"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"

	"github.com/m110/witchcraft/component"
	"github.com/m110/witchcraft/spell"
)

type Auras struct {
	query               *query.Query
	auraEffectResolvers []AuraEffectResolver
}

func NewAuras() *Auras {
	return &Auras{
		query: query.NewQuery(
			filter.Contains(
				component.AuraHolder,
			),
		),
		auraEffectResolvers: []AuraEffectResolver{
			ResolveAuraEffectNone,
			ResolveAuraEffectManaPercentRegen,
		},
	}
}

func (a *Auras) Update(w donburi.World) {
	a.query.Each(w, func(entry *donburi.Entry) {
		auraHolder := component.AuraHolder.Get(entry)

		aurasChanged := false

		for _, aura := range auraHolder.Auras {
			if aura.Timer != nil {
				aura.Timer.Update()

				if aura.Timer.IsReady() {
					aurasChanged = true
				}
			}

			if aura.TickTimer != nil {
				aura.TickTimer.Update()
				if aura.TickTimer.IsReady() {
					aura.TickTimer.Reset()

					resolved := false
					for _, resolver := range a.auraEffectResolvers {
						if resolver(entry, aura.Template) {
							resolved = true
							break
						}
					}

					if !resolved {
						panic(fmt.Sprintf("unknown aura effect: %v", aura.Template.Type))
					}
				}
			}
		}

		// Remove expired auras
		if aurasChanged {
			newAuras := make([]component.Aura, 0, len(auraHolder.Auras))

			for _, aura := range auraHolder.Auras {
				if aura.Timer == nil || !aura.Timer.IsReady() {
					newAuras = append(newAuras, aura)
				}
			}

			auraHolder.Auras = newAuras
		}

	})
}

type AuraEffectResolver func(caster *donburi.Entry, effect spell.AuraEffect) bool

func ResolveAuraEffectNone(caster *donburi.Entry, effect spell.AuraEffect) bool {
	if effect.Type != spell.AuraEffectTypeNone {
		return false
	}

	return true
}

func ResolveAuraEffectManaPercentRegen(caster *donburi.Entry, effect spell.AuraEffect) bool {
	if effect.Type != spell.AuraEffectTypeManaPercentRegen {
		return false
	}

	manaData := component.Mana.Get(caster)

	pct := float64(manaData.MaxMana) * float64(effect.Amount) / 100.0
	manaData.AddMana(int(pct))

	return true
}
