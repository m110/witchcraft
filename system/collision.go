package system

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"

	"github.com/m110/witchcraft/component"
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

		if collider.Collisions == nil {
			collider.Collisions = map[component.CollisionKey]component.Collision{}
		} else {
			for k, v := range collider.Collisions {
				v.Detected = false
				collider.Collisions[k] = v
			}
		}

		if collider.JustCollidedWith == nil {
			collider.JustCollidedWith = map[component.CollisionKey]struct{}{}
		} else {
			clear(collider.JustCollidedWith)
		}

		if collider.JustOutOfCollisionWith == nil {
			collider.JustOutOfCollisionWith = map[component.CollisionKey]struct{}{}
		} else {
			clear(collider.JustOutOfCollisionWith)
		}
	}

	for _, entry := range entries {
		collider := component.Collider.Get(entry)

		for _, other := range entries {
			// Can't collide with self
			if entry.Entity().Id() == other.Entity().Id() {
				continue
			}

			// Can't collide with destroyed entities
			if entry.HasComponent(component.Destroyed) || other.HasComponent(component.Destroyed) {
				continue
			}

			otherCollider := component.Collider.Get(other)

			col := component.CollisionKey{
				Layer: otherCollider.Layer,
				Other: other.Entity(),
			}

			if v, ok := collider.Collisions[col]; ok {
				// Collision already detected in this tick
				if v.Detected {
					continue
				}
			}

			// TODO The current approach doesn't take rotation into account
			rect := collider.Rect(entry)
			otherRect := otherCollider.Rect(other)

			if rect.Intersects(otherRect) {
				if v, ok := collider.Collisions[col]; ok {
					// Collision exists from previous ticks, and is still happening
					v.Detected = true
					if v.TimesSeen == 1 {
						v.TimesSeen++
					}

					collider.Collisions[col] = v
				} else {
					collider.Collisions[col] = component.Collision{
						Detected:  true,
						TimesSeen: 1,
					}
				}

				otherCol := component.CollisionKey{
					Layer: collider.Layer,
					Other: entry.Entity(),
				}

				if v, ok := otherCollider.Collisions[otherCol]; ok {
					// Collision exists from previous ticks, and is still happening
					v.Detected = true
					if v.TimesSeen == 1 {
						v.TimesSeen++
					}

					otherCollider.Collisions[otherCol] = v
				} else {
					otherCollider.Collisions[otherCol] = component.Collision{
						Detected:  true,
						TimesSeen: 1,
					}
				}
			}
		}
	}

	for _, entry := range entries {
		collider := component.Collider.Get(entry)

		for k, v := range collider.Collisions {
			if v.Detected {
				if v.TimesSeen == 1 {
					collider.JustCollidedWith[k] = struct{}{}
				}
			} else {
				collider.JustOutOfCollisionWith[k] = struct{}{}
				delete(collider.Collisions, k)
			}
		}
	}
}
