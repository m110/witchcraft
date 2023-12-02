package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type CharacterData struct {
	Body       Body
	Hair       *Hair
	FacialHair *Hair
	Equipment  Equipment
}

var Character = donburi.NewComponentType[CharacterData]()

func (c CharacterData) Image() *ebiten.Image {
	img := ebiten.NewImageFromImage(c.Body.Image)

	if c.Equipment.Feet != nil {
		img.DrawImage(c.Equipment.Feet.Image, nil)
	}

	if c.Equipment.Legs != nil {
		img.DrawImage(c.Equipment.Legs.Image, nil)
	}

	if c.Equipment.Chest != nil {
		img.DrawImage(c.Equipment.Chest.Image, nil)
	}

	if c.Hair != nil {
		img.DrawImage(c.Hair.Image, nil)
	}

	if c.FacialHair != nil {
		img.DrawImage(c.FacialHair.Image, nil)
	}

	if c.Equipment.Head != nil {
		img.DrawImage(c.Equipment.Head.Image, nil)
	}

	if c.Equipment.MainHand != nil {
		img.DrawImage(c.Equipment.MainHand.Image, nil)
	}

	if c.Equipment.OffHand != nil {
		img.DrawImage(c.Equipment.OffHand.Image, nil)
	}

	return img
}

type Equipment struct {
	Head     *Armor
	Chest    *Armor
	Legs     *Armor
	Feet     *Armor
	MainHand *Weapon
	OffHand  *Weapon
}

type Body struct {
	ID    int
	Index int
	Image *ebiten.Image
	Type  int
	Color int
}

type Hair struct {
	ID    int
	Index int
	Image *ebiten.Image
	Type  int
	Color int
}

type Armor struct {
	ID      int
	Index   int
	Image   *ebiten.Image
	Defense int
}

type Weapon struct {
	ID    int
	Index int
	Image *ebiten.Image
}
