package system

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"

	"github.com/m110/witchcraft/engine"

	"github.com/m110/witchcraft/archetype"
	"github.com/m110/witchcraft/component"
)

type DropExperience struct {
	query *query.Query
}

func NewDropExperience() *DropExperience {
	return &DropExperience{
		query: query.NewQuery(
			filter.Contains(
				component.Destroyed,
				component.DropExperience,
			),
		),
	}
}

func (s *DropExperience) Update(w donburi.World) {
	s.query.Each(w, func(e *donburi.Entry) {
		pos := transform.WorldPosition(e)
		exp := component.DropExperience.Get(e).Experience

		offset := 10.0
		for i := 0; i < exp; i++ {
			pos.X += engine.RandomFloatRange(-offset, offset)
			pos.Y += engine.RandomFloatRange(-offset, offset)
			archetype.NewExperiencePickup(w, pos)
		}
	})
}
