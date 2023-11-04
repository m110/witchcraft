package component

import "github.com/yohamta/donburi"

const (
	CollisionLayerPlayers ColliderLayer = iota
	CollisionLayerProjectiles
)

type ColliderLayer int

type ColliderData struct {
	Width      float64
	Height     float64
	Layer      ColliderLayer
	Collisions []Collision
}

type Collision struct {
	Layer ColliderLayer
	Other *donburi.Entry
}

var Collider = donburi.NewComponentType[ColliderData]()
