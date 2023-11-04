package component

import "github.com/yohamta/donburi"

type PlayerData struct {
	PlayerID int
}

var Player = donburi.NewComponentType[PlayerData]()
