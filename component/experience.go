package component

import (
	"github.com/yohamta/donburi"
)

type ExperienceData struct {
	Experience               int
	SkillPoints              int
	NextSkillPointExperience int
}

func (e *ExperienceData) AddExperience(amount int) {
	e.Experience += amount

	exp := e.Experience
	for exp >= e.NextSkillPointExperience {
		e.SkillPoints++
		exp -= e.NextSkillPointExperience
		e.NextSkillPointExperience = int(float64(e.NextSkillPointExperience) * 1.5)
	}

	e.Experience = exp
}

var Experience = donburi.NewComponentType[ExperienceData]()
