package spell

import (
	"time"

	"github.com/yohamta/donburi/features/math"

	"github.com/hajimehoshi/ebiten/v2"
)

type EffectType int

const (
	EffectTypeNone EffectType = iota
	EffectTypeSpawnProjectile
	EffectTypeApplyAura
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
	Image      *ebiten.Image
	Speed      float64
	Directions []math.Vec2
	Damage     int
	// TODO Perhaps "range" would be a better mechanic
	Duration     time.Duration
	OnHitEffects []Effect
}

type ApplyAuraData struct {
	AuraEffect AuraEffect
}

type SpawnEntityData struct {
	Type SpawnedEntityType
}
