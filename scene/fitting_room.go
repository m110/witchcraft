package scene

import (
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/samber/lo"
	"golang.org/x/image/colornames"

	"github.com/m110/witchcraft/assets"
	"github.com/m110/witchcraft/component"
)

const (
	fittingRoomOffsetX = 120
	fittingRoomOffsetY = 100

	fittingRoomTextOffsetX = 400
	fittingRoomTextOffsetY = 100
)

type FittingRoomPart struct {
	Name   string
	Index  *int
	Update func(delta int)
}

type FittingRoomIndexes struct {
	Hair       *int
	FacialHair *int
	Head       *int
	Chest      *int
	Legs       *int
	Feet       *int
	MainHand   *int
	OffHand    *int
}

type FittingRoom struct {
	context         Context
	character       *component.CharacterData
	image           *ebiten.Image
	parts           []FittingRoomPart
	activePartIndex int
}

func NewFittingRoom(context Context) *FittingRoom {
	f := &FittingRoom{
		context: context,
		character: &component.CharacterData{
			Body:       assets.Bodies[0],
			Hair:       nil,
			FacialHair: nil,
			Equipment:  component.Equipment{},
		},
	}

	f.onCharacterChanged()

	return f
}

func (f *FittingRoom) onCharacterChanged() {
	f.image = f.character.Image()

	indexes := FittingRoomIndexes{}
	if f.character.Hair != nil {
		indexes.Hair = &f.character.Hair.Index
	}
	if f.character.FacialHair != nil {
		indexes.FacialHair = &f.character.FacialHair.Index
	}
	if f.character.Equipment.Head != nil {
		indexes.Head = &f.character.Equipment.Head.Index
	}
	if f.character.Equipment.Chest != nil {
		indexes.Chest = &f.character.Equipment.Chest.Index
	}
	if f.character.Equipment.Legs != nil {
		indexes.Legs = &f.character.Equipment.Legs.Index
	}
	if f.character.Equipment.Feet != nil {
		indexes.Feet = &f.character.Equipment.Feet.Index
	}
	if f.character.Equipment.MainHand != nil {
		indexes.MainHand = &f.character.Equipment.MainHand.Index
	}
	if f.character.Equipment.OffHand != nil {
		indexes.OffHand = &f.character.Equipment.OffHand.Index
	}

	f.parts = []FittingRoomPart{
		{
			Name:  "Body",
			Index: &f.character.Body.Index,
			Update: func(delta int) {
				index := f.character.Body.Index + delta
				if index < 0 {
					index = len(assets.Bodies) - 1
				}
				if index >= len(assets.Bodies) {
					index = 0
				}
				f.character.Body = assets.Bodies[index]
			},
		},
		{
			Name:  "Hair",
			Index: indexes.Hair,
			Update: func(delta int) {
				next := nextIndex(indexes.Hair, delta, assets.Hairs)
				if next == nil {
					f.character.Hair = nil
				} else {
					f.character.Hair = &assets.Hairs[*next]
				}
			},
		},
		{
			Name:  "Facial Hair",
			Index: indexes.FacialHair,
			Update: func(delta int) {
				next := nextIndex(indexes.FacialHair, delta, assets.FacialHairs)
				if next == nil {
					f.character.FacialHair = nil
				} else {
					f.character.FacialHair = &assets.FacialHairs[*next]
				}
			},
		},
		{
			Name:  "Head",
			Index: indexes.Head,
			Update: func(delta int) {
				next := nextIndex(indexes.Head, delta, assets.HeadArmors)
				if next == nil {
					f.character.Equipment.Head = nil
				} else {
					f.character.Equipment.Head = &assets.HeadArmors[*next]
				}
			},
		},
		{
			Name:  "Chest",
			Index: indexes.Chest,
			Update: func(delta int) {
				next := nextIndex(indexes.Chest, delta, assets.ChestArmors)
				if next == nil {
					f.character.Equipment.Chest = nil
				} else {
					f.character.Equipment.Chest = &assets.ChestArmors[*next]
				}
			},
		},
		{
			Name:  "Legs",
			Index: indexes.Legs,
			Update: func(delta int) {
				next := nextIndex(indexes.Legs, delta, assets.LegsArmors)
				if next == nil {
					f.character.Equipment.Legs = nil
				} else {
					f.character.Equipment.Legs = &assets.LegsArmors[*next]
				}
			},
		},
		{
			Name:  "Feet",
			Index: indexes.Feet,
			Update: func(delta int) {
				next := nextIndex(indexes.Feet, delta, assets.FeetArmors)
				if next == nil {
					f.character.Equipment.Feet = nil
				} else {
					f.character.Equipment.Feet = &assets.FeetArmors[*next]
				}
			},
		},
		{
			Name:  "Main Hand",
			Index: indexes.MainHand,
			Update: func(delta int) {
				next := nextIndex(indexes.MainHand, delta, assets.MainHandWeapons)
				if next == nil {
					f.character.Equipment.MainHand = nil
				} else {
					f.character.Equipment.MainHand = &assets.MainHandWeapons[*next]
				}
			},
		},
		{
			Name:  "Off Hand",
			Index: indexes.OffHand,
			Update: func(delta int) {
				next := nextIndex(indexes.OffHand, delta, assets.OffHandWeapons)
				if next == nil {
					f.character.Equipment.OffHand = nil
				} else {
					f.character.Equipment.OffHand = &assets.OffHandWeapons[*next]
				}
			},
		},
	}
}

func (f *FittingRoom) Update() {
	for _, id := range ebiten.AppendGamepadIDs(nil) {
		if inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonRightRight) {
			f.context.SwitchToMainMenu()
		}

		if inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonLeftBottom) {
			f.activePartIndex++
			if f.activePartIndex >= len(f.parts) {
				f.activePartIndex = 0
			}
		}

		if inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonLeftTop) {
			f.activePartIndex--
			if f.activePartIndex < 0 {
				f.activePartIndex = len(f.parts) - 1
			}
		}

		if inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonLeftLeft) {
			f.parts[f.activePartIndex].Update(-1)
			f.onCharacterChanged()
		}

		if inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonLeftRight) {
			f.parts[f.activePartIndex].Update(1)
			f.onCharacterChanged()
		}
	}
}

func (f *FittingRoom) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(10, 10)
	op.GeoM.Translate(fittingRoomOffsetX, fittingRoomOffsetY)

	screen.DrawImage(f.image, op)

	for i, part := range f.parts {
		value := ""
		if part.Index == nil {
			value += "-"
		} else {
			value += strconv.Itoa(*part.Index)
		}

		color := colornames.White
		if i == f.activePartIndex {
			color = colornames.Yellow
		}

		text.Draw(screen, part.Name, assets.NarrowFont, fittingRoomTextOffsetX, fittingRoomTextOffsetY+i*30, color)
		text.Draw(screen, value, assets.NarrowFont, fittingRoomTextOffsetX+200, fittingRoomTextOffsetY+i*30, color)
	}
}

func nextIndex[T any](index *int, delta int, slice []T) *int {
	if index == nil {
		if delta > 0 {
			return lo.ToPtr(0)
		} else {
			return lo.ToPtr(len(slice) - 1)
		}
	}

	v := *index + delta

	if v < 0 || v >= len(slice) {
		return nil
	}

	return &v
}
