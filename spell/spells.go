package spell

import (
	stdmath "math"
	"time"

	"github.com/yohamta/donburi/features/math"

	"github.com/m110/witchcraft/assets"
)

type Spell struct {
	Name string

	IsChannel           bool
	ChannelTickDuration time.Duration
	ChannelTickManaCost int
	MaxChannelTime      time.Duration

	CastWhileMoving    bool
	ChannelWhileMoving bool

	ManaCost    int
	CastingTime time.Duration
	Cooldown    time.Duration

	OnCastEffects         []Effect
	OnChannelTickEffects  []Effect
	OnCastFinishedEffects []Effect
}

var FireBall, LightningBolt, Spark, ManaSurge, Quicksand, VenomBurst, ArcaneVolley, ArcaneBarrage Spell

func LoadSpells() {
	FireBall = Spell{
		Name:        "Fire Ball",
		ManaCost:    10,
		CastingTime: time.Millisecond * 200,
		Cooldown:    time.Second * 2,
		OnCastEffects: []Effect{
			{
				Type: EffectTypeSpawnProjectiles,
				Data: SpawnProjectilesData{
					Image:    assets.FireballProjectile,
					Speed:    5,
					Damage:   5,
					Duration: 0,
				},
			},
		},
	}
	LightningBolt = Spell{
		Name:        "Lightning Bolt",
		ManaCost:    25,
		CastingTime: 500 * time.Millisecond,
		Cooldown:    0,
		OnCastEffects: []Effect{
			{
				Type: EffectTypeSpawnProjectiles,
				Data: SpawnProjectilesData{
					Image:    assets.LightningBoltProjectile,
					Speed:    7,
					Damage:   2,
					Duration: time.Second * 2,
				},
			},
		},
	}
	Spark = Spell{
		Name:        "Spark",
		ManaCost:    5,
		CastingTime: 0,
		Cooldown:    time.Millisecond * 1500,
		OnCastEffects: []Effect{
			{
				Type: EffectTypeSpawnProjectiles,
				Data: SpawnProjectilesData{
					Image:    assets.SparkProjectile,
					Speed:    10,
					Damage:   1,
					Duration: time.Second * 1,
				},
			},
		},
	}
	ManaSurge = Spell{
		Name:        "Mana Surge",
		ManaCost:    0,
		CastingTime: 1 * time.Second,
		Cooldown:    30 * time.Second,
		OnCastEffects: []Effect{
			{
				Type: EffectTypeApplyAura,
				Data: ApplyAuraData{
					AuraEffect: AuraEffect{
						ID:       "mana-surge-regen",
						OnTick:   AuraEffectTypeManaPercentRegen,
						Image:    assets.IconManaSurge,
						Duration: 5 * time.Second,
						TickTime: 250 * time.Millisecond,
						Amount:   0.05,
					},
				},
			},
		},
	}
	Quicksand = Spell{
		Name:        "Quicksand",
		ManaCost:    20,
		CastingTime: 0,
		Cooldown:    10 * time.Second,
		OnCastEffects: []Effect{
			{
				Type: EffectTypeSpawnEntity,
				Data: SpawnEntityData{
					Type: SpawnedEntityTypeQuicksand,
				},
			},
		},
	}
	VenomBurst = Spell{
		Name:        "Venom Burst",
		ManaCost:    25,
		CastingTime: 100 * time.Millisecond,
		Cooldown:    1 * time.Second,
		OnCastEffects: []Effect{
			{
				Type: EffectTypeSpawnProjectiles,
				Data: SpawnProjectilesData{
					Image:    assets.VenomProjectile,
					Speed:    10,
					Damage:   1,
					Duration: 500 * time.Millisecond,
					Directions: func(direction math.Vec2) []math.Vec2 {
						return []math.Vec2{
							{X: 0, Y: -1},
							{X: 1, Y: -1},
							{X: 1, Y: 0},
							{X: 1, Y: 1},
							{X: 0, Y: 1},
							{X: -1, Y: 1},
							{X: -1, Y: 0},
							{X: -1, Y: -1},
						}
					},
					OnHitEffects: []Effect{
						{
							Type: EffectTypeApplyAura,
							Data: ApplyAuraData{
								AuraEffect: AuraEffect{
									ID: "venom-poison",
									// TODO Add proper icon
									Image:    nil,
									OnApply:  AuraEffectTypeNone,
									OnTick:   AuraEffectTypeDamage,
									Duration: 3 * time.Second,
									TickTime: 500 * time.Millisecond,
									Amount:   1,
								},
							},
						},
					},
				},
			},
		},
	}
	ArcaneVolley = Spell{
		Name:     "Arcane Volley",
		ManaCost: 10,
		Cooldown: 1 * time.Second,
		OnCastEffects: []Effect{
			{
				Type: EffectTypeSpawnProjectiles,
				Data: SpawnProjectilesData{
					Image:    assets.ArcaneProjectile,
					Speed:    7,
					Damage:   5,
					Duration: 2 * time.Second,
					Directions: func(direction math.Vec2) []math.Vec2 {
						baseAngle := stdmath.Atan2(direction.Y, direction.X)

						angles := []float64{
							15,
							10,
							0,
							-10,
							-15,
						}

						var directions []math.Vec2
						for _, angle := range angles {
							rad := angle * stdmath.Pi / 180.0
							directions = append(directions, math.Vec2{
								X: stdmath.Cos(baseAngle + rad),
								Y: stdmath.Sin(baseAngle + rad),
							})
						}

						return directions
					},
				},
			},
		},
	}
	ArcaneBarrage = Spell{
		Name:                "Arcane Barrage",
		IsChannel:           true,
		ChannelTickDuration: 250 * time.Millisecond,
		ChannelTickManaCost: 10,
		MaxChannelTime:      2 * time.Second,
		ManaCost:            5,
		CastingTime:         250 * time.Millisecond,
		Cooldown:            0,
		OnChannelTickEffects: []Effect{
			{
				Type: EffectTypeSpawnProjectiles,
				Data: SpawnProjectilesData{
					Image:    assets.ArcaneProjectile,
					Speed:    7,
					Damage:   5,
					Duration: 2 * time.Second,
					Directions: func(direction math.Vec2) []math.Vec2 {
						baseAngle := stdmath.Atan2(direction.Y, direction.X)

						angles := []float64{
							10,
							5,
							0,
							-5,
							-10,
						}

						var directions []math.Vec2
						for _, angle := range angles {
							rad := angle * stdmath.Pi / 180.0
							directions = append(directions, math.Vec2{
								X: stdmath.Cos(baseAngle + rad),
								Y: stdmath.Sin(baseAngle + rad),
							})
						}

						return directions
					},
				},
			},
		},
	}
}
