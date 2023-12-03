package system

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"

	"github.com/m110/witchcraft/component"
)

type AI struct {
	query *donburi.Query
}

func NewAI() *AI {
	return &AI{
		query: donburi.NewQuery(
			filter.Contains(
				component.AI,
			),
		),
	}
}

func (s *AI) Update(w donburi.World) {
	s.query.Each(w, func(entry *donburi.Entry) {
		ai := component.AI.Get(entry)

		switch ai.Type {
		case component.AITypeNone:
		case component.AITypeMoveTowardsTarget:
			target := component.Targeter.Get(entry).Target
			if target == nil {
				return
			}

			moveSpeed := component.Mover.Get(entry).Speed
			entryPos := transform.WorldPosition(entry)
			targetPos := transform.WorldPosition(target)

			delta := targetPos.Sub(entryPos).Normalized().MulScalar(moveSpeed)

			component.Velocity.Get(entry).Velocity = delta
		}
	})
}
