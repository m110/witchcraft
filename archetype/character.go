package archetype

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"

	"github.com/m110/cliche-rpg/assets"
	"github.com/m110/cliche-rpg/component"
)

func NewRandomCharacter(w donburi.World, position math.Vec2) {
	c := w.Entry(
		w.Create(
			transform.Transform,
			component.Sprite,
			component.Character,
		),
	)

	transform.Transform.Get(c).LocalPosition = position
	transform.Transform.Get(c).LocalScale = math.Vec2{X: 2, Y: 2}

	character := component.CharacterData{
		Body:       assets.RandomFrom(assets.Bodies),
		Hair:       assets.RandomFromOrEmpty(assets.Hairs),
		FacialHair: assets.RandomFromOrEmpty(assets.FacialHairs),
		Equipment: component.Equipment{
			Head:     assets.RandomFromOrEmpty(assets.HeadArmors),
			Chest:    assets.RandomFromOrEmpty(assets.ChestArmors),
			Legs:     assets.RandomFromOrEmpty(assets.LegsArmors),
			Feet:     assets.RandomFromOrEmpty(assets.FeetArmors),
			MainHand: assets.RandomFromOrEmpty(assets.MainHandWeapons),
			OffHand:  assets.RandomFromOrEmpty(assets.OffHandWeapons),
		},
	}

	component.Character.Set(c, &character)

	component.Sprite.SetValue(c, component.SpriteData{
		Image: character.Image(),
	})
}
