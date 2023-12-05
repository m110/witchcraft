package archetype

import (
	"time"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"

	"github.com/m110/witchcraft/assets"
	"github.com/m110/witchcraft/component"
	"github.com/m110/witchcraft/engine"
	"github.com/m110/witchcraft/spell"
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
		Layer: component.SpriteLayerFloorUnits,
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
			component.DropExperience,
		),
	)

	c := component.CharacterData{
		Body: component.AllBodyParts.Bodies[6],
	}

	transform.Transform.Get(o).LocalScale = math.Vec2{X: 2, Y: 2}

	component.Sprite.SetValue(o, component.SpriteData{
		Image: c.Image(),
		Layer: component.SpriteLayerUnits,
	})

	component.Collider.SetValue(o, component.ColliderData{
		Width:  float64(c.Image().Bounds().Dx()),
		Height: float64(c.Image().Bounds().Dy()),
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
		Speed: 0.75,
	})

	component.DropExperience.SetValue(o, component.DropExperienceData{
		Experience: engine.RandomIntRange(1, 3),
	})

	return o
}

func NewQuicksand(w donburi.World, summoner *donburi.Entry, teamID component.TeamID) *donburi.Entry {
	q := w.Entry(
		w.Create(
			transform.Transform,
			component.Sprite,
			component.Collider,
			component.Team,
			component.TimeToLive,
			component.AuraEmitter,
			component.Summon,
		),
	)

	component.Sprite.SetValue(q, component.SpriteData{
		Image: assets.QuicksandArea,
		Pivot: component.SpritePivotCenter,
		Layer: component.SpriteLayerFloorEffect,
	})

	component.Collider.SetValue(q, component.ColliderData{
		Offset: math.Vec2{
			X: -50,
			Y: -50,
		},
		Width:  100,
		Height: 100,
		Layer:  component.CollisionLayerEffects,
	})

	component.Team.SetValue(q, component.TeamData{
		TeamID: teamID,
	})

	component.TimeToLive.Set(q, &component.TimeToLiveData{
		Timer: engine.NewTimer(time.Second * 5),
	})

	component.AuraEmitter.SetValue(q, component.AuraEmitterData{
		AuraEffect: spell.AuraEffect{
			ID:      "quicksand-slow",
			OnApply: spell.AuraEffectTypeSlowMovement,
			Image:   assets.IconSlow,
			Amount:  0.5,
		},
	})

	component.Summon.SetValue(q, component.SummonData{
		Summoner: summoner,
	})

	return q
}
