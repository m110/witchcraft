package system

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"

	"github.com/m110/witchcraft/component"
	"github.com/m110/witchcraft/engine"
)

type Spawner struct {
	query *donburi.Query
}

func NewSpawner() *Spawner {
	return &Spawner{
		query: donburi.NewQuery(
			filter.Contains(
				component.Spawner,
			),
		),
	}
}

func (s *Spawner) Update(w donburi.World) {
	s.query.Each(w, func(entry *donburi.Entry) {
		spawner := component.Spawner.Get(entry)

		if spawner.Timer == nil {
			spawner.Timer = engine.NewTimer(spawner.Interval.Random())
			return
		}

		spawner.Timer.Update()

		if spawner.Timer.IsReady() {
			spawner.Timer.Reset()
			spawner.Timer.SetTarget(spawner.Interval.Random())

			pos := transform.WorldPosition(entry)
			offsetX := engine.RandomFloatRange(-10, 10)
			offsetY := engine.RandomFloatRange(-10, 10)

			pos.X += offsetX
			pos.Y += offsetY

			for i := 0; i < spawner.Count.Random(); i++ {
				entity := spawner.SpawnFunc(w)
				transform.GetTransform(entity).LocalPosition = pos
			}
		}
	})
}
