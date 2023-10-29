package system

import (
	"math"

	math2 "github.com/yohamta/donburi/features/math"

	"github.com/m110/witchcraft/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type Crosshair struct {
	query *query.Query
}

func NewCrosshair() *Crosshair {
	return &Crosshair{
		query: query.NewQuery(
			filter.Contains(component.Crosshair),
		),
	}
}

func (c *Crosshair) Update(w donburi.World) {
	c.query.Each(w, func(entry *donburi.Entry) {
		parent, ok := transform.GetParent(entry)
		if !ok {
			return
		}

		direction := component.Direction.Get(parent)

		const distance = 30

		angle := math.Atan2(direction.Direction.Y, direction.Direction.X)
		transform.GetTransform(entry).LocalPosition = math2.Vec2{
			X: distance * math.Cos(angle),
			Y: distance * math.Sin(angle),
		}
	})
}
