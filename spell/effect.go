package spell

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type EffectType int

const (
	EffectTypeNone EffectType = iota
	EffectTypeSpawnProjectile
	EffectTypeApplyAuraOnCaster
	EffectTypeSpawnEntity
)

type SpawnedEntityType int

const (
	SpawnedEntityTypeNone SpawnedEntityType = iota
	SpawnedEntityTypeQuicksand
)

type Effect struct {
	Type EffectType
	Data any
}

type SpawnProjectileData struct {
	Image    *ebiten.Image
	Speed    float64
	Damage   int
	Duration time.Duration
}

type ApplyAuraData struct {
	AuraTemplate AuraEffect
}

type SpawnEntityData struct {
	Type SpawnedEntityType
}
