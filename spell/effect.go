package spell

import "github.com/yohamta/donburi"

type Effect interface {
	Resolve(caster *donburi.Entry, w donburi.World)
}

type SpawnProjectileEffect struct {
}

func (e SpawnProjectileEffect) Resolve(caster *donburi.Entry, world donburi.World) {
	// Spawn projectile
}
