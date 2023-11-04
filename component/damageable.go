package component

import "github.com/yohamta/donburi"

type DamageableData struct {
	Team   int
	Damage int
}

var Damageable = donburi.NewComponentType[DamageableData]()
