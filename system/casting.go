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

var spellEffectResolvers = []SpellEffectResolver{
	ResolveSpellEffectNone,
	ResolveSpellEffectSpawnProjectile,
	ResolveSpellEffectApplyAura,
	ResolveSpellEffectSpawnEntity,
}

type Casting struct {
	query *query.Query
	debug *component.DebugData
}

func NewCasting() *Casting {
	return &Casting{
		query: query.NewQuery(
			filter.Contains(
				component.Caster,
			),
		),
	}
}

func (c *Casting) Update(w donburi.World) {
	if c.debug == nil {
		debug, ok := query.NewQuery(filter.Contains(component.Debug)).First(w)
		if ok {
			c.debug = component.Debug.Get(debug)
		}
	}

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

			manaCost := preparedSpell.Template.ManaCost
			if c.debug.Enabled {
				manaCost = 0
			}

			mana := component.Mana.Get(entry)
			if !mana.UseMana(manaCost) {
				// Not enough mana — can't cast
				return
			}
		}

		// Casting in progress
		preparedSpell.CastingTimer.Update()

		// The casting is done — cast the spell
		if preparedSpell.CastingTimer.IsReady() {
			resolveSpellEffects(entry, entry, preparedSpell.Template.OnCastEffects)

			preparedSpell.CastingTimer.Reset()

			if !c.debug.Enabled {
				preparedSpell.CooldownTimer.Reset()
			}
		}
	})
}

func resolveSpellEffects(caster *donburi.Entry, target *donburi.Entry, effects []spell.Effect) {
	for _, effect := range effects {
		resolved := false
		for _, resolver := range spellEffectResolvers {
			if resolver(caster, target, effect) {
				resolved = true
				break
			}
		}

		if !resolved {
			panic(fmt.Sprintf("unknown spell effect: %v", effect.Type))
		}
	}
}

type SpellEffectResolver func(caster *donburi.Entry, target *donburi.Entry, effect spell.Effect) bool

func ResolveSpellEffectNone(caster *donburi.Entry, target *donburi.Entry, effect spell.Effect) bool {
	if effect.Type != spell.EffectTypeNone {
		return false
	}

	return true
}

func ResolveSpellEffectSpawnProjectile(caster *donburi.Entry, target *donburi.Entry, effect spell.Effect) bool {
	if effect.Type != spell.EffectTypeSpawnProjectiles {
		return false
	}

	data := effect.Data.(spell.SpawnProjectilesData)
	archetype.NewProjectiles(target, data)

	return true
}

func ResolveSpellEffectApplyAura(caster *donburi.Entry, target *donburi.Entry, effect spell.Effect) bool {
	if effect.Type != spell.EffectTypeApplyAura {
		return false
	}

	data := effect.Data.(spell.ApplyAuraData)
	aura := component.NewAura(caster, data.AuraEffect)

	applyAura(target, aura)

	return true
}

func ResolveSpellEffectSpawnEntity(caster *donburi.Entry, target *donburi.Entry, effect spell.Effect) bool {
	if effect.Type != spell.EffectTypeSpawnEntity {
		return false
	}

	data := effect.Data.(spell.SpawnEntityData)

	switch data.Type {
	case spell.SpawnedEntityTypeNone:
	case spell.SpawnedEntityTypeQuicksand:
		teamID := component.Team.Get(caster).TeamID
		q := archetype.NewQuicksand(caster.World, target, teamID)
		transform.GetTransform(q).LocalPosition = transform.WorldPosition(target)
	}

	return true
}
