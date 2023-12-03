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

		for _, collision := range collider.Collisions {
			if collision.Other.HasComponent(component.Health) {
				if collision.Other.HasComponent(component.Team) {
					otherTeam := component.Team.Get(collision.Other)
					if team.TeamID == otherTeam.TeamID {
						continue
					}
				}

				damageEntity(collision.Other, damageable.Damage)
				entry.Remove()
				break
			}
		}
	})
}

func damageEntity(entry *donburi.Entry, damage int) {
	health := component.Health.Get(entry)
	health.Damage(damage)

	if health.Health <= 0 {
		entry.Remove()
	}

	archetype.NewDamageText(entry.World, damage, transform.WorldPosition(entry))
}
