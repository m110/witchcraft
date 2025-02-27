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

var FireBall, LightningBolt, Spark, ManaSurge, Quicksand, VenomBurst, ArcaneVolley, ArcaneBarrage, WaterBeam, FrostNova, MeteorShower, ArcaneMissiles, ShadowStrike, BlinkDash, PoisonDagger Spell

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
		Name:            "Venom Burst",
		ManaCost:        25,
		CastingTime:     500 * time.Millisecond,
		CastWhileMoving: true,
		Cooldown:        1 * time.Second,
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
	WaterBeam = Spell{
		Name:                "Water Beam",
		ManaCost:            5,
		CastingTime:         1 * time.Second,
		Cooldown:            1 * time.Second,
		IsChannel:           true,
		ChannelTickDuration: 10 * time.Millisecond,
		ChannelTickManaCost: 1,
		MaxChannelTime:      2 * time.Second,
		ChannelWhileMoving:  true,
		OnChannelTickEffects: []Effect{
			{
				Type: EffectTypeSpawnProjectiles,
				Data: SpawnProjectilesData{
					Image:    assets.WaterProjectile,
					Speed:    3,
					Damage:   1,
					Duration: 1 * time.Second,
				},
			},
		},
	}

	FrostNova = Spell{
		Name:        "Frost Nova",
		ManaCost:    30,
		CastingTime: 750 * time.Millisecond,
		Cooldown:    8 * time.Second,
		OnCastEffects: []Effect{
			{
				Type: EffectTypeSpawnProjectiles,
				Data: SpawnProjectilesData{
					Image:    assets.FrostProjectile,
					Speed:    6,
					Damage:   3,
					Duration: 1 * time.Second,
					Directions: func(direction math.Vec2) []math.Vec2 {
						var directions []math.Vec2
						for i := 0; i < 8; i++ {
							angle := float64(i) * 45 * stdmath.Pi / 180.0
							directions = append(directions, math.Vec2{
								X: stdmath.Cos(angle),
								Y: stdmath.Sin(angle),
							})
						}
						return directions
					},
					OnHitEffects: []Effect{
						{
							Type: EffectTypeApplyAura,
							Data: ApplyAuraData{
								AuraEffect: AuraEffect{
									ID:       "frost-slow",
									Image:    assets.IconSlow,
									OnApply:  AuraEffectTypeSlowMovement,
									OnTick:   AuraEffectTypeNone,
									Duration: 3 * time.Second,
									Amount:   0.5, // 50% slow
								},
							},
						},
					},
				},
			},
		},
	}

	MeteorShower = Spell{
		Name:        "Meteor Shower",
		ManaCost:    40,
		CastingTime: 1 * time.Second,
		Cooldown:    12 * time.Second,
		OnCastEffects: []Effect{
			{
				Type: EffectTypeSpawnProjectiles,
				Data: SpawnProjectilesData{
					Image:    assets.MeteorProjectile,
					Speed:    4,
					Damage:   15,
					Duration: 3 * time.Second,
					Directions: func(direction math.Vec2) []math.Vec2 {
						// Spawn 3 meteors in a line pattern
						baseAngle := stdmath.Atan2(direction.Y, direction.X)
						var directions []math.Vec2
						directions = append(directions, math.Vec2{
							X: stdmath.Cos(baseAngle),
							Y: stdmath.Sin(baseAngle),
						})
						return directions
					},
				},
			},
		},
	}

	ArcaneMissiles = Spell{
		Name:        "Arcane Missiles",
		ManaCost:    20,
		CastingTime: 300 * time.Millisecond,
		Cooldown:    3 * time.Second,
		OnCastEffects: []Effect{
			{
				Type: EffectTypeSpawnProjectiles,
				Data: SpawnProjectilesData{
					Image:    assets.ArcaneProjectile,
					Speed:    8,
					Damage:   2,
					Duration: 2 * time.Second,
					Directions: func(direction math.Vec2) []math.Vec2 {
						baseAngle := stdmath.Atan2(direction.Y, direction.X)

						// Generate 3 homing missiles
						var directions []math.Vec2
						for i := 0; i < 3; i++ {
							// Slight random offset per missile
							randOffset := (float64(i) - 1.0) * 0.1
							directions = append(directions, math.Vec2{
								X: stdmath.Cos(baseAngle) + randOffset,
								Y: stdmath.Sin(baseAngle) + randOffset,
							})
						}
						return directions
					},
				},
			},
		},
	}

	ShadowStrike = Spell{
		Name:            "Shadow Strike",
		ManaCost:        15,
		CastingTime:     200 * time.Millisecond,
		CastWhileMoving: true,
		Cooldown:        2 * time.Second,
		OnCastEffects: []Effect{
			{
				Type: EffectTypeSpawnProjectiles,
				Data: SpawnProjectilesData{
					Image:    assets.ShadowProjectile,
					Speed:    12,
					Damage:   8,
					Duration: 500 * time.Millisecond,
					Directions: func(direction math.Vec2) []math.Vec2 {
						baseAngle := stdmath.Atan2(direction.Y, direction.X)

						var directions []math.Vec2
						// Main direction
						directions = append(directions, math.Vec2{
							X: stdmath.Cos(baseAngle),
							Y: stdmath.Sin(baseAngle),
						})
						return directions
					},
					OnHitEffects: []Effect{
						{
							Type: EffectTypeApplyAura,
							Data: ApplyAuraData{
								AuraEffect: AuraEffect{
									ID:       "shadow-strike-slow",
									Image:    assets.IconSlow,
									OnApply:  AuraEffectTypeSlowMovement,
									OnTick:   AuraEffectTypeNone,
									TickTime: 500 * time.Millisecond,
									Duration: 2 * time.Second,
									Amount:   0.3, // 30% slow
								},
							},
						},
					},
				},
			},
		},
	}

	BlinkDash = Spell{
		Name:            "Blink Dash",
		ManaCost:        10,
		CastingTime:     0,
		CastWhileMoving: true,
		Cooldown:        3 * time.Second,
		OnCastEffects: []Effect{
			{
				Type: EffectTypeApplyAura,
				Data: ApplyAuraData{
					AuraEffect: AuraEffect{
						ID:       "blink-speed",
						OnApply:  AuraEffectTypeSpeedBoost,
						OnTick:   AuraEffectTypeNone,
						Image:    assets.IconSlow, // Reuse existing icon
						TickTime: 100 * time.Millisecond,
						Duration: 500 * time.Millisecond,
						Amount:   3.0, // 3x speed boost
					},
				},
			},
		},
	}

	PoisonDagger = Spell{
		Name:        "Poison Dagger",
		ManaCost:    20,
		CastingTime: 250 * time.Millisecond,
		Cooldown:    5 * time.Second,
		OnCastEffects: []Effect{
			{
				Type: EffectTypeSpawnProjectiles,
				Data: SpawnProjectilesData{
					Image:    assets.PoisonProjectile,
					Speed:    9,
					Damage:   3,
					Duration: 1 * time.Second,
					Directions: func(direction math.Vec2) []math.Vec2 {
						baseAngle := stdmath.Atan2(direction.Y, direction.X)

						// Throw three daggers in a narrow spread
						angles := []float64{
							5,
							0,
							-5,
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
					OnHitEffects: []Effect{
						{
							Type: EffectTypeApplyAura,
							Data: ApplyAuraData{
								AuraEffect: AuraEffect{
									ID:       "poison-damage",
									OnApply:  AuraEffectTypeNone,
									OnTick:   AuraEffectTypeDamage,
									Image:    assets.IconSlow, // Reuse existing icon
									Duration: 4 * time.Second,
									TickTime: 500 * time.Millisecond,
									Amount:   2, // 2 damage per tick
								},
							},
						},
					},
				},
			},
		},
	}
}
