package system

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"

	"github.com/m110/witchcraft/component"
)

type CollisionExperiencePickups struct {
	query *query.Query
}

func NewCollisionExperiencePickups() *CollisionExperiencePickups {
	return &CollisionExperiencePickups{
		query: query.NewQuery(
			filter.Contains(
				component.ExperiencePickup,
				component.Collider,
			),
		),
	}
}

func (c *CollisionExperiencePickups) Update(w donburi.World) {
	c.query.Each(w, func(entry *donburi.Entry) {
		collider := component.Collider.Get(entry)

		for collision := range collider.JustCollidedWith {
			other := w.Entry(collision.Other)

			if !other.HasComponent(component.Collector) {
				continue
			}

			var experience *component.ExperienceData
			if other.HasComponent(component.Experience) {
				experience = component.Experience.Get(other)
			} else {
				parent, ok := transform.GetParent(other)
				if !ok {
					panic("collector has no experience and no parent")
				}

				experience = component.Experience.Get(parent)
			}

			pickup := component.ExperiencePickup.Get(entry)
			experience.AddExperience(pickup.Amount)

			component.Destroy(entry)
		}
	})
}
