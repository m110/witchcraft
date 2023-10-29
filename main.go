package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/m110/witchcraft/assets"
	"github.com/m110/witchcraft/scene"
)

var (
	screenWidth  = 800
	screenHeight = 600
)

type Scene interface {
	Update()
	Draw(screen *ebiten.Image)
}

type Game struct {
	scene Scene
}

func NewGame() *Game {
	assets.MustLoadAssets()

	g := &Game{}
	g.switchToWitchcraft()
	return g
}

func (g *Game) switchToCharacterSelect() {
	g.scene = scene.NewCharacterSelect(screenWidth, screenHeight)
}

func (g *Game) switchToWitchcraft() {
	g.scene = scene.NewWitchcraft(screenWidth, screenHeight)
}

func (g *Game) Update() error {
	g.scene.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.scene.Draw(screen)
}

func (g *Game) Layout(width, height int) (int, int) {
	return width, height
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	rand.Seed(time.Now().UTC().UnixNano())

	err := ebiten.RunGame(NewGame())
	if err != nil {
		log.Fatal(err)
	}
}
