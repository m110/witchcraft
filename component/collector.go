package component

import "github.com/yohamta/donburi"

type CollectorData struct {
}

var Collector = donburi.NewComponentType[CollectorData]()
