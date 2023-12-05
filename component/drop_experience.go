package component

import "github.com/yohamta/donburi"

type DropExperienceData struct {
	Experience int
}

var DropExperience = donburi.NewComponentType[DropExperienceData]()
