package scene

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"

	"github.com/m110/witchcraft/archetype"
	"github.com/m110/witchcraft/component"
	"github.com/m110/witchcraft/system"
)

type JoinedPlayer struct {
	GamePadID ebiten.GamepadID
	Class     archetype.Class
}

type Battle struct {
	world     donburi.World
	systems   []System
	drawables []Drawable

	joinedPlayers []JoinedPlayer

	screenWidth  int
	screenHeight int
}

func NewBattle(screenWidth int, screenHeight int, joinedPlayers []JoinedPlayer) *Battle {
	g := &Battle{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,

		joinedPlayers: joinedPlayers,
	}

	fmt.Println("Joined players:", joinedPlayers)

	g.loadLevel()

	return g
}

func (w *Battle) loadLevel() {
	render := system.NewRenderer()

	w.systems = []System{
		system.NewVelocity(),
		system.NewControls(),
		system.NewCasting(),
		system.NewProgressBar(),
		system.NewMana(),
		system.NewCrosshair(),
		render,
		system.NewTimeToLive(),
	}

	w.drawables = []Drawable{
		render,
		system.NewText(),
	}

	w.world = w.createWorld()
	w.spawnCharacters()
}

func (w *Battle) createWorld() donburi.World {
	world := donburi.NewWorld()

	archetype.NewCamera(world, math.Vec2{})

	game := world.Entry(world.Create(component.Game))
	donburi.SetValue(game, component.Game, component.GameData{
		Settings: component.Settings{
			ScreenWidth:  w.screenWidth,
			ScreenHeight: w.screenHeight,
		},
	})

	world.Create(component.Debug)

	return world
}

func (w *Battle) Update() {
	for _, s := range w.systems {
		s.Update(w.world)
	}
}

func (w *Battle) spawnCharacters() {
	offset := 48.0

	archetype.NewCharacter(w.world, math.Vec2{X: offset, Y: offset})
}

func (w *Battle) Draw(screen *ebiten.Image) {
	for _, s := range w.drawables {
		s.Draw(w.world, screen)
	}
}
