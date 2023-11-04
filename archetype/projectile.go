package archetype

import (
	stdmath "math"

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
			component.Damageable,
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

	dir := component.Direction.Get(caster).Direction

	angle := stdmath.Atan2(dir.Y, dir.X) * 180.0 / stdmath.Pi
	transform.GetTransform(projectile).LocalRotation = angle

	magnitude := stdmath.Sqrt(dir.X*dir.X + dir.Y*dir.Y)
	normalizedX := dir.X / magnitude
	normalizedY := dir.Y / magnitude

	component.Velocity.Get(projectile).Velocity = math.Vec2{
		X: normalizedX * data.Speed,
		Y: normalizedY * data.Speed,
	}

	player := component.Player.Get(caster)

	component.Damageable.Set(projectile, &component.DamageableData{
		Team:   player.TeamID,
		Damage: data.Damage,
	})

	return projectile
}
