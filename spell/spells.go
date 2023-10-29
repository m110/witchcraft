package spell

import "time"

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

var (
	FireBall = Spell{
		Name:        "Fire Ball",
		ManaCost:    10,
		CastingTime: time.Millisecond * 500,
		Cooldown:    time.Second * 2,
		OnCastEffects: []Effect{
			SpawnProjectileEffect{},
		},
	}
	LightningBolt = Spell{
		Name:        "Lightning Bolt",
		ManaCost:    15,
		CastingTime: time.Second,
		Cooldown:    0,
		OnCastEffects: []Effect{
			SpawnProjectileEffect{},
		},
	}
	Spark = Spell{
		Name:        "Spark",
		ManaCost:    5,
		CastingTime: 0,
		Cooldown:    time.Millisecond * 1500,
		OnCastEffects: []Effect{
			SpawnProjectileEffect{},
		},
	}
)
