package system

import (
	"github.com/m110/witchcraft/component"
	"github.com/m110/witchcraft/engine"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type Collision struct {
	query *query.Query
}

func NewCollision() *Collision {
	return &Collision{
		query: query.NewQuery(
			filter.Contains(
				transform.Transform,
				component.Collider,
			),
		),
	}
}

func (c *Collision) Update(w donburi.World) {
	var entries []*donburi.Entry
	c.query.Each(w, func(entry *donburi.Entry) {
		entries = append(entries, entry)
	})

	for _, entry := range entries {
		collider := component.Collider.Get(entry)

		for _, other := range entries {
			if entry.Entity().Id() == other.Entity().Id() {
				continue
			}

			otherCollider := component.Collider.Get(other)

			pos := transform.WorldPosition(entry)
			otherPos := transform.WorldPosition(other)

			// TODO The current approach doesn't take rotation into account
			// TODO The current approach doesn't take scale into account
			rect := engine.NewRect(pos.X, pos.Y, collider.Width, collider.Height)
			otherRect := engine.NewRect(otherPos.X, otherPos.Y, otherCollider.Width, otherCollider.Height)

			if rect.Intersects(otherRect) {
				collider.Collisions = append(collider.Collisions, component.Collision{
					Layer: otherCollider.Layer,
					Other: other,
				})

				otherCollider.Collisions = append(otherCollider.Collisions, component.Collision{
					Layer: collider.Layer,
					Other: entry,
				})
			}
		}
	}
}
