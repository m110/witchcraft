package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"

	"github.com/m110/witchcraft/assets"
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
}

type Hair struct {
	ID    int
	Index int
	Image *ebiten.Image
}

type Armor struct {
	ID    int
	Index int
	Image *ebiten.Image
}

type Weapon struct {
	ID    int
	Index int
	Image *ebiten.Image
}

type BodyParts struct {
	Bodies          []Body
	Hairs           []Hair
	FacialHairs     []Hair
	HeadArmors      []Armor
	ChestArmors     []Armor
	LegsArmors      []Armor
	FeetArmors      []Armor
	MainHandWeapons []Weapon
	OffHandWeapons  []Weapon
}

var AllBodyParts BodyParts

func LoadBodyParts() {
	bp := BodyParts{}

	for i, b := range assets.Bodies {
		bp.Bodies = append(bp.Bodies, Body{
			ID:    b.ID,
			Index: i,
			Image: b.Image,
		})
	}

	bp.Hairs = bodyPartToHairSlice(assets.Hairs)
	bp.FacialHairs = bodyPartToHairSlice(assets.FacialHairs)
	bp.HeadArmors = bodyPartToArmorSlice(assets.HeadArmors)
	bp.ChestArmors = bodyPartToArmorSlice(assets.ChestArmors)
	bp.LegsArmors = bodyPartToArmorSlice(assets.LegsArmors)
	bp.FeetArmors = bodyPartToArmorSlice(assets.FeetArmors)
	bp.MainHandWeapons = bodyPartToWeaponSlice(assets.MainHandWeapons)
	bp.OffHandWeapons = bodyPartToWeaponSlice(assets.OffHandWeapons)

	AllBodyParts = bp
}

func bodyPartToHairSlice(bodyParts []assets.BodyPart) []Hair {
	hairs := make([]Hair, len(bodyParts))

	for i, h := range bodyParts {
		hairs[i] = Hair{
			ID:    h.ID,
			Index: i,
			Image: h.Image,
		}
	}

	return hairs
}

func bodyPartToArmorSlice(bodyParts []assets.BodyPart) []Armor {
	armors := make([]Armor, len(bodyParts))

	for i, a := range bodyParts {
		armors[i] = Armor{
			ID:    a.ID,
			Index: i,
			Image: a.Image,
		}
	}

	return armors
}

func bodyPartToWeaponSlice(bodyParts []assets.BodyPart) []Weapon {
	weapons := make([]Weapon, len(bodyParts))

	for i, w := range bodyParts {
		weapons[i] = Weapon{
			ID:    w.ID,
			Index: i,
			Image: w.Image,
		}
	}

	return weapons
}
