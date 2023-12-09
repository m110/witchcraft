package archetype

import (
	stdmath "math"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"

	"github.com/m110/witchcraft/component"
	"github.com/m110/witchcraft/engine"
	"github.com/m110/witchcraft/spell"
)

func NewProjectiles(caster *donburi.Entry, data spell.SpawnProjectilesData) []*donburi.Entry {
	w := caster.World

	var projectiles []*donburi.Entry

	dir := component.Direction.Get(caster).Direction

	var directions []math.Vec2
	if data.Directions == nil {
		directions = []math.Vec2{dir}
	} else {
		directions = data.Directions(dir)
	}

	for _, dir := range directions {
		projectile := w.Entry(
			w.Create(
				transform.Transform,
				component.Sprite,
				component.Velocity,
				component.Collider,
				component.Damageable,
				component.Team,
				component.SpellEffect,
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

		bounds := data.Image.Bounds()
		component.Collider.SetValue(projectile, component.ColliderData{
			Offset: math.Vec2{
				X: -float64(bounds.Dx()) / 2,
				Y: -float64(bounds.Dy()) / 2,
			},
			Width:  float64(bounds.Dx()),
			Height: float64(bounds.Dy()),
			Layer:  component.CollisionLayerProjectiles,
		})

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
			Damage: data.Damage,
		})

		component.Team.Set(projectile, &component.TeamData{
			TeamID: component.TeamID(player.PlayerID),
		})

		component.SpellEffect.Set(projectile, &component.SpellEffectData{
			Caster:  caster,
			Effects: data.OnHitEffects,
		})

		projectiles = append(projectiles, projectile)
	}

	return projectiles
}
