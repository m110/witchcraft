package archetype

import (
	"time"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"

	"github.com/m110/witchcraft/assets"
	"github.com/m110/witchcraft/component"
	"github.com/m110/witchcraft/engine"
)

func NewSpawner(w donburi.World) *donburi.Entry {
	s := w.Entry(
		w.Create(
			transform.Transform,
			component.Sprite,
			component.Spawner,
		),
	)

	component.Sprite.SetValue(s, component.SpriteData{
		Image: assets.Spawner,
	})

	component.Spawner.SetValue(s, component.SpawnerData{
		Count: engine.IntRange{
			Min: 1,
			Max: 3,
		},
		Interval: engine.DurationRange{
			Min: time.Second,
			Max: 5 * time.Second,
		},
		SpawnFunc: NewOrc,
	})

	return s
}

func NewOrc(w donburi.World) *donburi.Entry {
	o := w.Entry(
		w.Create(
			transform.Transform,
			component.Sprite,
			component.Collider,
			component.Health,
			component.AuraHolder,
			component.AI,
			component.Targeter,
			component.Seeker,
			component.Mover,
			component.Velocity,
		),
	)

	c := component.CharacterData{
		Body: component.AllBodyParts.Bodies[6],
	}

	transform.Transform.Get(o).LocalScale = math.Vec2{X: 2, Y: 2}

	component.Sprite.SetValue(o, component.SpriteData{
		Image: c.Image(),
	})

	component.Collider.SetValue(o, component.ColliderData{
		Width:  float64(c.Image().Bounds().Dx() * 2),
		Height: float64(c.Image().Bounds().Dy() * 2),
		Layer:  component.CollisionLayerEnemies,
	})

	component.Health.SetValue(o, component.HealthData{
		Health:    10,
		MaxHealth: 10,
	})

	component.Seeker.SetValue(o, component.SeekerData{
		Type: component.SeekerTypeNearestPlayer,
	})

	component.AI.SetValue(o, component.AIData{
		Type: component.AITypeMoveTowardsTarget,
	})

	component.Mover.SetValue(o, component.MoverData{
		Speed: 1,
	})

	return o
}
