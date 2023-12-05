package component

import "github.com/yohamta/donburi"

type ExperiencePickupData struct {
	Amount int
}

var ExperiencePickup = donburi.NewComponentType[ExperiencePickupData]()
