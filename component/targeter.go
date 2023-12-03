package component

import "github.com/yohamta/donburi"

type TargetData struct {
	Target *donburi.Entry
}

var Targeter = donburi.NewComponentType[TargetData]()
