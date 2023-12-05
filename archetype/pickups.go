package archetype

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"

	"github.com/m110/witchcraft/assets"
	"github.com/m110/witchcraft/component"
)

func NewExperiencePickup(w donburi.World, pos math.Vec2) *donburi.Entry {
	pickup := w.Entry(
		w.Create(
			transform.Transform,
			component.Sprite,
			component.Collider,
			component.ExperiencePickup,
		),
	)

	transform.GetTransform(pickup).LocalPosition = pos

	component.Sprite.SetValue(pickup, component.SpriteData{
		Image: assets.Experience,
		Layer: component.SpriteLayerPickups,
	})

	bounds := component.Sprite.Get(pickup).Image.Bounds()

	component.Collider.SetValue(pickup, component.ColliderData{
		Width:  float64(bounds.Dx()),
		Height: float64(bounds.Dy()),
		Layer:  component.CollisionLayerPickups,
	})

	component.ExperiencePickup.SetValue(pickup, component.ExperiencePickupData{
		Amount: 1,
	})

	return pickup
}
