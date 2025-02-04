package component

import (
	"github.com/yohamta/donburi"

	"github.com/m110/witchcraft/engine"
)

type TimeToLiveData struct {
	Timer *engine.Timer
}

var TimeToLive = donburi.NewComponentType[TimeToLiveData]()
