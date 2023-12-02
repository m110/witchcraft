package spell

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type AuraEffectType int

const (
	AuraEffectTypeNone AuraEffectType = iota
	AuraEffectTypeManaPercentRegen
)

type AuraEffect struct {
	Type AuraEffectType

	// TODO Perhaps could follow the "data" pattern here
	Image    *ebiten.Image
	Duration time.Duration
	TickTime time.Duration
	Amount   int
}
