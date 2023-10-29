package component

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
)

type DirectionData struct {
	Direction math.Vec2
}

var Direction = donburi.NewComponentType[DirectionData]()
