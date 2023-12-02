package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/m110/witchcraft/assets"
	"github.com/m110/witchcraft/component"
	"github.com/m110/witchcraft/scene"
	"github.com/m110/witchcraft/spell"
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

func NewGame(startScene string) *Game {
	assets.MustLoadAssets()
	component.LoadBodyParts()
	spell.LoadSpells()

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

	switch startScene {
	case "":
		g.switchToMainMenu()
	case "fitting_room":
		g.switchToFittingRoom()
	case "character_select":
		g.switchToCharacterSelect()
	default:
		panic("unknown start scene")
	}

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

	startScene := os.Getenv("START_SCENE")

	err := ebiten.RunGame(NewGame(startScene))
	if err != nil {
		log.Fatal(err)
	}
}
