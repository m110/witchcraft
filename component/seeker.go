package component

import "github.com/yohamta/donburi"

type SeekerType int

const (
	SeekerTypeNone SeekerType = iota
	SeekerTypeNearestPlayer
)

type SeekerData struct {
	Type SeekerType
}

var Seeker = donburi.NewComponentType[SeekerData]()
