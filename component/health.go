package component

import "github.com/yohamta/donburi"

type HealthData struct {
	Health    int
	MaxHealth int
}

func (d *HealthData) Damage(damage int) {
	d.Health -= damage
	if d.Health <= 0 {
		d.Health = 0
	}
}

var Health = donburi.NewComponentType[HealthData]()
