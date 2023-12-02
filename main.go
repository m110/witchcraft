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
	sceneContext scene.Context
	scene        Scene
	mainMenu     *scene.MainMenu
}

func NewGame() *Game {
	assets.MustLoadAssets()

	g := &Game{}

	g.sceneContext = scene.Context{
		ScreenWidth:             screenWidth,
		ScreenHeight:            screenHeight,
		SwitchToMainMenu:        g.switchToMainMenu,
		SwitchToFittingRoom:     g.switchToFittingRoom,
		SwitchToBattle:          g.switchToBattle,
		SwitchToCharacterSelect: g.switchToCharacterSelect,
	}

	// Keeping the main menu scene in memory to remember the active item index
	g.mainMenu = scene.NewMainMenu(g.sceneContext)

	g.switchToMainMenu()

	return g
}

func (g *Game) switchToMainMenu() {
	g.scene = g.mainMenu
}

func (g *Game) switchToFittingRoom() {
	g.scene = scene.NewFittingRoom(g.sceneContext)
}

func (g *Game) switchToCharacterSelect() {
	g.scene = scene.NewCharacterSelect(g.sceneContext)
}

func (g *Game) switchToBattle(joinedPlayers []scene.JoinedPlayer) {
	g.scene = scene.NewBattle(g.sceneContext, joinedPlayers)
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
