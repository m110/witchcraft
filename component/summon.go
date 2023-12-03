package component

import "github.com/yohamta/donburi"

type SummonData struct {
	Summoner *donburi.Entry
}

var Summon = donburi.NewComponentType[SummonData]()
