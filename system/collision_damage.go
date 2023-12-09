package system

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"

	"github.com/m110/witchcraft/archetype"
	"github.com/m110/witchcraft/component"
)

type CollisionDamage struct {
	query *donburi.Query
}

func NewCollisionDamage() *CollisionDamage {
	return &CollisionDamage{
		query: donburi.NewQuery(
			filter.Contains(
				component.Collider,
				component.Team,
				component.Damageable,
			),
		),
	}
}

func (c *CollisionDamage) Update(w donburi.World) {
	c.query.Each(w, func(entry *donburi.Entry) {
		collider := component.Collider.Get(entry)
		damageable := component.Damageable.Get(entry)

		team := component.Team.Get(entry)

		for collision := range collider.Collisions {
			other := w.Entry(collision.Other)
			if other.HasComponent(component.Health) {
				if other.HasComponent(component.Team) {
					otherTeam := component.Team.Get(other)
					if team.TeamID == otherTeam.TeamID {
						continue
					}
				}

				// TODO It seems like sometimes it triggers twice and sometimes doesn't trigger at all
				damageEntity(other, damageable.Damage)
				entry.AddComponent(component.Destroyed)

				if entry.HasComponent(component.SpellEffect) {
					se := component.SpellEffect.Get(entry)
					resolveSpellEffects(se.Caster, other, se.Effects)
				}

				break
			}
		}
	})
}

func damageEntity(entry *donburi.Entry, damage int) {
	health := component.Health.Get(entry)
	health.Damage(damage)

	if health.Health <= 0 {
		component.Destroy(entry)
	}

	archetype.NewDamageText(entry.World, damage, transform.WorldPosition(entry))
}
