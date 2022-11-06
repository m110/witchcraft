package component

import (
	"github.com/yohamta/donburi"
)

type CameraData struct {
}

var Camera = donburi.NewComponentType[CameraData]()
