package system

import (
	"fmt"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"

	"github.com/m110/witchcraft/archetype"
	"github.com/m110/witchcraft/component"
	"github.com/m110/witchcraft/spell"
)

type Casting struct {
	query                *query.Query
	spellEffectResolvers []SpellEffectResolver
}

func NewCasting() *Casting {
	return &Casting{
		query: query.NewQuery(
			filter.Contains(
				component.Caster,
			),
		),
		spellEffectResolvers: []SpellEffectResolver{
			ResolveSpellEffectNone,
			ResolveSpellEffectSpawnProjectile,
			ResolveSpellEffectApplyAuraOnCaster,
			ResolveSpellEffectSpawnEntity,
		},
	}
}

func (c *Casting) Update(w donburi.World) {
	c.query.Each(w, func(entry *donburi.Entry) {
		caster := component.Caster.Get(entry)

		// Update cooldown for all spells, except the prepared one
		for i := range caster.KnownSpells {
			if caster.PreparedSpellIndex == nil || i != *caster.PreparedSpellIndex {
				caster.KnownSpells[i].CooldownTimer.Update()
			}
		}

		preparedSpell, ok := caster.PreparedSpell()
		if !ok {
			return
		}

		if !caster.IsCasting {
			if preparedSpell.CastingTimer.IsStarted() {
				// If the casting has been started and then stopped, interrupt the cast
				preparedSpell.CastingTimer.Reset()
			} else {
				// Spell prepared, but not casting at the moment, update cooldown
				preparedSpell.CooldownTimer.Update()
			}

			return
		}

		// Casting just starting
		if !preparedSpell.CastingTimer.IsStarted() {
			preparedSpell.CooldownTimer.Update()

			if !preparedSpell.CooldownTimer.IsReady() {
				// Cooldown is not ready — can't cast
				return
			}

			mana := component.Mana.Get(entry)
			if !mana.UseMana(preparedSpell.Template.ManaCost) {
				// Not enough mana — can't cast
				return
			}
		}

		// Casting in progress
		preparedSpell.CastingTimer.Update()

		// The casting is done — cast the spell
		if preparedSpell.CastingTimer.IsReady() {
			for _, effect := range preparedSpell.Template.OnCastEffects {
				resolved := false
				for _, resolver := range c.spellEffectResolvers {
					if resolver(entry, effect) {
						resolved = true
						break
					}
				}

				if !resolved {
					panic(fmt.Sprintf("unknown spell effect: %v", effect.Type))
				}
			}

			preparedSpell.CastingTimer.Reset()
			preparedSpell.CooldownTimer.Reset()
		}
	})
}

type SpellEffectResolver func(caster *donburi.Entry, effect spell.Effect) bool

func ResolveSpellEffectNone(caster *donburi.Entry, effect spell.Effect) bool {
	if effect.Type != spell.EffectTypeNone {
		return false
	}

	return true
}

func ResolveSpellEffectSpawnProjectile(caster *donburi.Entry, effect spell.Effect) bool {
	if effect.Type != spell.EffectTypeSpawnProjectile {
		return false
	}

	data := effect.Data.(spell.SpawnProjectileData)
	archetype.NewProjectile(caster, data)

	return true
}

func ResolveSpellEffectApplyAuraOnCaster(caster *donburi.Entry, effect spell.Effect) bool {
	if effect.Type != spell.EffectTypeApplyAuraOnCaster {
		return false
	}

	data := effect.Data.(spell.ApplyAuraData)
	aura := component.NewAura(caster, data.AuraTemplate)

	applyAura(caster, aura)

	return true
}

func ResolveSpellEffectSpawnEntity(caster *donburi.Entry, effect spell.Effect) bool {
	if effect.Type != spell.EffectTypeSpawnEntity {
		return false
	}

	data := effect.Data.(spell.SpawnEntityData)

	switch data.Type {
	case spell.SpawnedEntityTypeNone:
	case spell.SpawnedEntityTypeQuicksand:
		teamID := component.Team.Get(caster).TeamID
		q := archetype.NewQuicksand(caster.World, caster, teamID)
		transform.GetTransform(q).LocalPosition = transform.WorldPosition(caster)
	}

	return true
}
