package system

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"

	"github.com/m110/cliche-rpg/component"
)

type Casting struct {
	query *query.Query
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
	c.query.EachEntity(w, func(entry *donburi.Entry) {
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
				effect.Resolve(entry, w)
			}

			preparedSpell.CastingTimer.Reset()
			preparedSpell.CooldownTimer.Reset()
		}
	})
}
