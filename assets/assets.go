package assets

import (
	"bytes"
	"embed"
	"encoding/xml"
	"image"
	_ "image/png"
	"io/fs"
	"math/rand"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/m110/witchcraft/component"
)

var (
	//go:embed fonts/kenney-future.ttf
	normalFontData []byte
	//go:embed fonts/kenney-future-narrow.ttf
	narrowFontData []byte

	//go:embed *
	assetsFS embed.FS

	SmallFont  font.Face
	NormalFont font.Face
	NarrowFont font.Face

	Bodies          []component.Body
	Hairs           []component.Hair
	FacialHairs     []component.Hair
	HeadArmors      []component.Armor
	ChestArmors     []component.Armor
	LegsArmors      []component.Armor
	FeetArmors      []component.Armor
	MainHandWeapons []component.Weapon
	OffHandWeapons  []component.Weapon
)

const (
	bodyClass       = "body"
	hairClass       = "hair"
	facialHairClass = "facial_hair"
	armorClass      = "armor"
	weaponClass     = "weapon"

	slotProperty = "slot"
	headSlot     = "head"
	chestSlot    = "chest"
	legsSlot     = "legs"
	feetSlot     = "feet"

	mainHandSlot = "main_hand"
	offHandSlot  = "off_hand"
)

func MustLoadAssets() {
	tileSetContent, err := assetsFS.ReadFile("levels/characters.tsx")
	if err != nil {
		panic(err)
	}

	var tileSet tiled.Tileset
	err = xml.Unmarshal(tileSetContent, &tileSet)
	if err != nil {
		panic(err)
	}

	tileSetImageFile, err := fs.ReadFile(assetsFS, filepath.Join("levels", tileSet.Image.Source))
	if err != nil {
		panic(err)
	}

	tileSetImage := mustNewEbitenImage(tileSetImageFile)

	for _, t := range tileSet.Tiles {
		img := mustSubImage(tileSetImage, tileSet, t.ID)

		switch t.Class {
		case bodyClass:
			Bodies = append(Bodies, component.Body{
				ID:    int(t.ID),
				Index: len(Bodies),
				Image: img,
			})
		case hairClass:
			Hairs = append(Hairs, component.Hair{
				ID:    int(t.ID),
				Index: len(Hairs),
				Image: img,
			})
		case facialHairClass:
			FacialHairs = append(FacialHairs, component.Hair{
				ID:    int(t.ID),
				Index: len(FacialHairs),
				Image: img,
			})
		case armorClass:
			switch t.Properties.GetString(slotProperty) {
			case headSlot:
				HeadArmors = append(HeadArmors, component.Armor{
					ID:    int(t.ID),
					Index: len(HeadArmors),
					Image: img,
				})
			case chestSlot:
				ChestArmors = append(ChestArmors, component.Armor{
					ID:    int(t.ID),
					Index: len(ChestArmors),
					Image: img,
				})
			case legsSlot:
				LegsArmors = append(LegsArmors, component.Armor{
					ID:    int(t.ID),
					Index: len(LegsArmors),
					Image: img,
				})
			case feetSlot:
				FeetArmors = append(FeetArmors, component.Armor{
					ID:    int(t.ID),
					Index: len(FeetArmors),
					Image: img,
				})
			}
		case weaponClass:
			switch t.Properties.GetString(slotProperty) {
			case mainHandSlot:
				MainHandWeapons = append(MainHandWeapons, component.Weapon{
					ID:    int(t.ID),
					Index: len(MainHandWeapons),
					Image: img,
				})
			case offHandSlot:
				OffHandWeapons = append(OffHandWeapons, component.Weapon{
					ID:    int(t.ID),
					Index: len(OffHandWeapons),
					Image: img,
				})
			}
		}
	}

	SmallFont = mustLoadFont(normalFontData, 10)
	NormalFont = mustLoadFont(normalFontData, 24)
	NarrowFont = mustLoadFont(narrowFontData, 24)
}

func mustSubImage(tileSetImage *ebiten.Image, ts tiled.Tileset, id uint32) *ebiten.Image {
	width := ts.TileWidth
	height := ts.TileHeight

	col := int(id) % ts.Columns
	row := int(id) / ts.Columns

	// Plus one because of 1px margin
	if col > 0 {
		width += 1
	}
	if row > 0 {
		height += 1
	}

	sx := col * width
	sy := row * height

	return tileSetImage.SubImage(
		image.Rect(sx, sy, sx+ts.TileWidth, sy+ts.TileHeight),
	).(*ebiten.Image)
}

func mustLoadFont(data []byte, size int) font.Face {
	f, err := opentype.Parse(data)
	if err != nil {
		panic(err)
	}

	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    float64(size),
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		panic(err)
	}

	return face
}

func mustNewEbitenImage(data []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

func RandomFrom[T comparable](list []T) T {
	index := rand.Intn(len(list))
	return list[index]
}

func RandomFromOrEmpty[T comparable](list []T) *T {
	index := rand.Intn(len(list) + 1)
	if index == len(list) {
		return nil
	}
	return &list[index]
}
