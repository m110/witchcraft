package component

import "github.com/yohamta/donburi"

type TeamID int

type TeamData struct {
	TeamID TeamID
}

var Team = donburi.NewComponentType[TeamData]()
