package component

import (
	"github.com/yohamta/donburi"

	"github.com/m110/witchcraft/spell"
)

type SpellEffectData struct {
	Caster  *donburi.Entry
	Effects []spell.Effect
}

var SpellEffect = donburi.NewComponentType[SpellEffectData]()
