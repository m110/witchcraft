package spell

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"

	"github.com/m110/witchcraft/assets"
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

var FireBall, LightningBolt, Spark, ManaSurge Spell

func LoadSpells() {
	fireballImage.Fill(colornames.Red)
	lightningBoltImage.Fill(colornames.Lightblue)
	sparkImage.Fill(colornames.Lightyellow)

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
		ManaCost:    25,
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
	ManaSurge = Spell{
		Name:        "Mana Surge",
		Runes:       []Rune{},
		ManaCost:    0,
		CastingTime: 1 * time.Second,
		Cooldown:    30 * time.Second,
		OnCastEffects: []Effect{
			{
				Type: EffectTypeApplyAuraOnCaster,
				Data: ApplyAuraData{
					AuraTemplate: AuraEffect{
						Type:     AuraEffectTypeManaPercentRegen,
						Image:    assets.IconManaSurge,
						Duration: 5 * time.Second,
						TickTime: 250 * time.Millisecond,
						Amount:   5,
					},
				},
			},
		},
	}
}
