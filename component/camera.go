package component

import (
	"github.com/yohamta/donburi"
)

type CameraData struct {
}

var Camera = donburi.NewComponentType[CameraData]()

func GetCamera(entry *donburi.Entry) *CameraData {
	return donburi.Get[CameraData](entry, Camera)
}
