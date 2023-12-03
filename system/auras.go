package system

import (
	"fmt"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"

	"github.com/m110/witchcraft/archetype"
	"github.com/m110/witchcraft/component"
	"github.com/m110/witchcraft/spell"
)

var auraEffectOnApplyResolvers = []AuraEffectResolver{
	ResolveAuraEffectNone,
	ResolveAuraEffectSlowMovement,
}
var auraEffectOnTickResolvers = []AuraEffectResolver{
	ResolveAuraEffectNone,
	ResolveAuraEffectManaPercentRegen,
}

type Auras struct {
	query *query.Query
}

func NewAuras() *Auras {
	return &Auras{
		query: query.NewQuery(
			filter.Contains(
				component.AuraHolder,
			),
		),
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
					for _, resolver := range auraEffectOnTickResolvers {
						if resolver(entry, aura.Effect.OnTick, aura.Effect) {
							resolved = true
							break
						}
					}

					if !resolved {
						panic(fmt.Sprintf("unknown aura effect on tick: %v", aura.Effect.OnTick))
					}
				}
			}
		}

		// Remove expired auras
		if aurasChanged {
			newAuras := make([]component.Aura, 0, len(auraHolder.Auras))

			for i, aura := range auraHolder.Auras {
				if aura.Timer != nil && aura.Timer.IsReady() {
					onAuraRemoved(entry, i)
				} else {
					newAuras = append(newAuras, aura)
				}
			}

			auraHolder.Auras = newAuras
		}

	})
}

type AuraEffectResolver func(caster *donburi.Entry, auraEffectType spell.AuraEffectType, effect spell.AuraEffect) bool

func ResolveAuraEffectNone(caster *donburi.Entry, auraEffectType spell.AuraEffectType, effect spell.AuraEffect) bool {
	if auraEffectType != spell.AuraEffectTypeNone {
		return false
	}

	return true
}

func ResolveAuraEffectManaPercentRegen(target *donburi.Entry, auraEffectType spell.AuraEffectType, effect spell.AuraEffect) bool {
	if auraEffectType != spell.AuraEffectTypeManaPercentRegen {
		return false
	}

	manaData := component.Mana.Get(target)

	pct := float64(manaData.MaxMana) * effect.Amount
	manaData.AddMana(int(pct))

	return true
}

func ResolveAuraEffectSlowMovement(caster *donburi.Entry, auraEffectType spell.AuraEffectType, effect spell.AuraEffect) bool {
	if auraEffectType != spell.AuraEffectTypeSlowMovement {
		return false
	}

	return true
}

func applyAura(target *donburi.Entry, aura component.Aura) {
	// TODO How to make sure this logic is in one place?
	ah := component.AuraHolder.Get(target)
	ah.ApplyAura(aura)

	resolved := false
	for _, r := range auraEffectOnApplyResolvers {
		if r(target, aura.Effect.OnApply, aura.Effect) {
			resolved = true
			break
		}
	}

	if !resolved {
		panic(fmt.Sprintf("unknown aura effect on apply: %v", aura.Effect.OnApply))
	}

	if ah.UI != nil {
		icon := archetype.NewAuraIcon(target.World, aura)
		component.List.Get(ah.UI).Append(ah.UI, icon)
	}
}

func onAuraRemoved(target *donburi.Entry, index int) {
	ah := component.AuraHolder.Get(target)

	if ah.UI != nil {
		component.List.Get(ah.UI).Remove(index)
	}
}
