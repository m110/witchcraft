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

		team := component.Team.Get(entry)

		for _, collision := range collider.Collisions {
			if collision.Other.HasComponent(component.AuraHolder) {
				if collision.Other.HasComponent(component.Team) {
					otherTeam := component.Team.Get(collision.Other)
					if team.TeamID == otherTeam.TeamID {
						// TODO remove
						// continue
					}
				}

				aura := component.NewAura(entry, emitter.AuraTemplate)
				holder := component.AuraHolder.Get(collision.Other)

				exists := false
				for _, a := range holder.Auras {
					if a.Equals(aura) {
						exists = true
						break
					}
				}

				if !exists {
					applyAura(collision.Other, aura)
				}
			}
		}
	})
}
