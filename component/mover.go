package component

import "github.com/yohamta/donburi"

type MoverData struct {
	Speed float64
}

var Mover = donburi.NewComponentType[MoverData]()
