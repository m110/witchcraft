package spell

import (
	"time"

	"golang.org/x/image/colornames"

	"github.com/hajimehoshi/ebiten/v2"
)

type Rune int

const (
	RuneAether Rune = iota
	RuneChaos
	RuneVoid
	RuneDream
)

type Spell struct {
	Name  string
	Runes []Rune

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
	fireballImage      = ebiten.NewImage(10, 10)
	lightningBoltImage = ebiten.NewImage(15, 3)
	sparkImage         = ebiten.NewImage(5, 2)
)

func init() {
	fireballImage.Fill(colornames.Red)
	lightningBoltImage.Fill(colornames.Lightblue)
	sparkImage.Fill(colornames.Lightyellow)
}

var (
	FireBall = Spell{
		Name:        "Fire Ball",
		Runes:       []Rune{RuneAether, RuneAether},
		ManaCost:    10,
		CastingTime: time.Millisecond * 500,
		Cooldown:    time.Second * 2,
		OnCastEffects: []Effect{
			{
				Type: EffectTypeSpawnProjectile,
				Data: SpawnProjectileData{
					Image:    fireballImage,
					Speed:    5,
					Damage:   5,
					Duration: 0,
				},
			},
		},
	}
	LightningBolt = Spell{
		Name:        "Lightning Bolt",
		Runes:       []Rune{RuneAether, RuneChaos, RuneVoid},
		ManaCost:    15,
		CastingTime: time.Second,
		Cooldown:    0,
		OnCastEffects: []Effect{
			{
				Type: EffectTypeSpawnProjectile,
				Data: SpawnProjectileData{
					Image:    lightningBoltImage,
					Speed:    7,
					Damage:   2,
					Duration: time.Second * 2,
				},
			},
		},
	}
	Spark = Spell{
		Name:        "Spark",
		Runes:       []Rune{RuneVoid, RuneDream, RuneVoid},
		ManaCost:    5,
		CastingTime: 0,
		Cooldown:    time.Millisecond * 1500,
		OnCastEffects: []Effect{
			{
				Type: EffectTypeSpawnProjectile,
				Data: SpawnProjectileData{
					Image:    sparkImage,
					Speed:    10,
					Damage:   1,
					Duration: time.Second * 1,
				},
			},
		},
	}
)
