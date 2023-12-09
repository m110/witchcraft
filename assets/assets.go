package assets

import (
	"bytes"
	"embed"
	"encoding/xml"
	"image"
	_ "image/png"
	"io/fs"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
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

	Bodies          []BodyPart
	Hairs           []BodyPart
	FacialHairs     []BodyPart
	HeadArmors      []BodyPart
	ChestArmors     []BodyPart
	LegsArmors      []BodyPart
	FeetArmors      []BodyPart
	MainHandWeapons []BodyPart
	OffHandWeapons  []BodyPart

	IconManaSurge *ebiten.Image
	IconSlow      *ebiten.Image

	Spawner *ebiten.Image

	FireballProjectile      *ebiten.Image
	LightningBoltProjectile *ebiten.Image
	SparkProjectile         = ebiten.NewImage(5, 2)
	VenomProjectile         = ebiten.NewImage(20, 5)
	ArcaneProjectile        = ebiten.NewImage(5, 5)
	QuicksandArea           = ebiten.NewImage(100, 100)

	Experience *ebiten.Image
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

type BodyPart struct {
	ID    int
	Image *ebiten.Image
}

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
			Bodies = append(Bodies, BodyPart{
				ID:    int(t.ID),
				Image: img,
			})
		case hairClass:
			Hairs = append(Hairs, BodyPart{
				ID:    int(t.ID),
				Image: img,
			})
		case facialHairClass:
			FacialHairs = append(FacialHairs, BodyPart{
				ID:    int(t.ID),
				Image: img,
			})
		case armorClass:
			switch t.Properties.GetString(slotProperty) {
			case headSlot:
				HeadArmors = append(HeadArmors, BodyPart{
					ID:    int(t.ID),
					Image: img,
				})
			case chestSlot:
				ChestArmors = append(ChestArmors, BodyPart{
					ID:    int(t.ID),
					Image: img,
				})
			case legsSlot:
				LegsArmors = append(LegsArmors, BodyPart{
					ID:    int(t.ID),
					Image: img,
				})
			case feetSlot:
				FeetArmors = append(FeetArmors, BodyPart{
					ID:    int(t.ID),
					Image: img,
				})
			}
		case weaponClass:
			switch t.Properties.GetString(slotProperty) {
			case mainHandSlot:
				MainHandWeapons = append(MainHandWeapons, BodyPart{
					ID:    int(t.ID),
					Image: img,
				})
			case offHandSlot:
				OffHandWeapons = append(OffHandWeapons, BodyPart{
					ID:    int(t.ID),
					Image: img,
				})
			}
		}
	}

	SmallFont = mustLoadFont(normalFontData, 10)
	NormalFont = mustLoadFont(normalFontData, 24)
	NarrowFont = mustLoadFont(narrowFontData, 24)

	IconManaSurge = mustLoadImage("icons/mana-surge.png")
	IconSlow = mustLoadImage("icons/slow.png")

	Spawner = mustLoadImage("entities/spawner.png")

	FireballProjectile = mustLoadImage("spells/fireball.png")
	LightningBoltProjectile = mustLoadImage("spells/lightning-bolt.png")

	SparkProjectile.Fill(colornames.Lightyellow)
	VenomProjectile.Fill(colornames.Limegreen)
	ArcaneProjectile.Fill(colornames.Blueviolet)
	QuicksandArea.Fill(colornames.Sandybrown)

	Experience = mustLoadImage("pickups/experience.png")
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

func mustLoadImage(path string) *ebiten.Image {
	data, err := fs.ReadFile(assetsFS, path)
	if err != nil {
		panic(err)
	}

	return mustNewEbitenImage(data)
}

func mustNewEbitenImage(data []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}
