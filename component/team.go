package component

import "github.com/yohamta/donburi"

type TeamData struct {
	TeamID int
}

var Team = donburi.NewComponentType[TeamData]()
