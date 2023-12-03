package component

import "github.com/yohamta/donburi"

type AIType int

const (
	AITypeNone AIType = iota
	AITypeMoveTowardsTarget
)

type AIData struct {
	Type AIType
}

var AI = donburi.NewComponentType[AIData]()
