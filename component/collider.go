package component

import "github.com/yohamta/donburi"

const (
	CollisionLayerPlayers ColliderLayer = iota
	CollisionLayerProjectiles
	CollisionLayerEffects
	CollisionLayerEnemies
)

type ColliderLayer int

type ColliderData struct {
	Width  float64
	Height float64
	Layer  ColliderLayer

	Collisions             map[CollisionKey]Collision
	JustCollidedWith       map[CollisionKey]struct{}
	JustOutOfCollisionWith map[CollisionKey]struct{}
}

type CollisionKey struct {
	Layer ColliderLayer
	Other donburi.Entity
}

type Collision struct {
	TimesSeen int
	Detected  bool
}

var Collider = donburi.NewComponentType[ColliderData]()
