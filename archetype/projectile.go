package archetype

import (
	"github.com/m110/witchcraft/engine"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"

	"github.com/m110/witchcraft/component"
	"github.com/m110/witchcraft/spell"
)

func NewProjectile(caster *donburi.Entry, data spell.SpawnProjectileData) *donburi.Entry {
	w := caster.World

	projectile := w.Entry(
		w.Create(
			transform.Transform,
			component.Sprite,
			component.Velocity,
		),
	)

	if data.Duration > 0 {
		projectile.AddComponent(component.TimeToLive)
		component.TimeToLive.Set(projectile, &component.TimeToLiveData{
			Timer: engine.NewTimer(data.Duration),
		})
	}

	casterPos := transform.WorldPosition(caster)
	transform.SetWorldPosition(projectile, casterPos)

	component.Sprite.Set(projectile, &component.SpriteData{
		Image: data.Image,
		Layer: component.SpriteLayerProjectiles,
		Pivot: component.SpritePivotCenter,
	})

	component.Velocity.Get(projectile).Velocity = math.Vec2{
		X: data.Speed,
		Y: data.Speed,
	}

	return projectile
}
