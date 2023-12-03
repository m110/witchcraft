package component

import (
	"github.com/yohamta/donburi"

	"github.com/m110/witchcraft/spell"
)

type AuraEmitterData struct {
	AuraTemplate spell.AuraEffect
}

var AuraEmitter = donburi.NewComponentType[AuraEmitterData]()
