package system

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"

	"github.com/m110/witchcraft/component"
)

type CollisionApplyAura struct {
	query *donburi.Query
}

func NewCollisionApplyAura() *CollisionApplyAura {
	return &CollisionApplyAura{
		query: donburi.NewQuery(
			filter.Contains(
				component.Collider,
				component.Team,
				component.AuraEmitter,
			),
		),
	}
}

func (c *CollisionApplyAura) Update(w donburi.World) {
	c.query.Each(w, func(entry *donburi.Entry) {
		collider := component.Collider.Get(entry)
		emitter := component.AuraEmitter.Get(entry)

		var source *donburi.Entry
		if entry.HasComponent(component.Summon) {
			source = component.Summon.Get(entry).Summoner
		} else {
			source = entry
		}

		team := component.Team.Get(entry)

		for collision := range collider.JustCollidedWith {
			other := w.Entry(collision.Other)

			if other.HasComponent(component.AuraHolder) {
				if other.HasComponent(component.Team) {
					otherTeam := component.Team.Get(other)
					if team.TeamID == otherTeam.TeamID {
						// TODO remove
						// continue
					}
				}

				aura := component.NewAura(source, emitter.AuraEffect)
				holder := component.AuraHolder.Get(other)

				exists := false
				for _, a := range holder.Auras {
					if a.Equals(aura) {
						exists = true
						break
					}
				}

				if !exists {
					applyAura(other, aura)
				}
			}
		}

		for collision := range collider.JustOutOfCollisionWith {
			other := w.Entry(collision.Other)

			if other.HasComponent(component.AuraHolder) {
				holder := component.AuraHolder.Get(other)

				for i, a := range holder.Auras {
					if a.EqualsTo(source, emitter.AuraEffect) {
						holder.Auras = append(holder.Auras[:i], holder.Auras[i+1:]...)
						onAuraRemoved(other, i)
						break
					}
				}
			}
		}
	})
}
