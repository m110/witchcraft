package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type InputData struct {
	Disabled bool

	MoveUpKey    ebiten.Key
	MoveRightKey ebiten.Key
	MoveDownKey  ebiten.Key
	MoveLeftKey  ebiten.Key

	SpellKeyA ebiten.Key
	SpellKeyB ebiten.Key
	SpellKeyC ebiten.Key

	CastKey ebiten.Key
}

var Input = donburi.NewComponentType[InputData]()
