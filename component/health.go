package component

import "github.com/yohamta/donburi"

type HealthData struct {
	Health    int
	MaxHealth int
}

func (d *HealthData) Damage() {
	if d.Health <= 0 {
		return
	}

	d.Health--
}

var Health = donburi.NewComponentType[HealthData]()
