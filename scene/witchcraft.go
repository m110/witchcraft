package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"

	"github.com/m110/cliche-rpg/archetype"
	"github.com/m110/cliche-rpg/component"
	"github.com/m110/cliche-rpg/system"
)

type Witchcraft struct {
	world     donburi.World
	systems   []System
	drawables []Drawable

	screenWidth  int
	screenHeight int
}

func NewWitchcraft(screenWidth int, screenHeight int) *Witchcraft {
	g := &Witchcraft{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
	}

	g.loadLevel()

	return g
}

func (w *Witchcraft) loadLevel() {
	render := system.NewRenderer()

	w.systems = []System{
		render,
	}

	w.drawables = []Drawable{
		render,
	}

	w.world = w.createWorld()
	w.spawnCharacters()
}

func (w *Witchcraft) createWorld() donburi.World {
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

func (w *Witchcraft) Update() {
	for _, s := range w.systems {
		s.Update(w.world)
	}
}

func (w *Witchcraft) spawnCharacters() {
	offset := 48.0

	for i := 0; i < 6; i++ {
		archetype.NewCharacter(w.world, math.Vec2{X: offset + float64(i)*offset, Y: offset})
	}
}

func (w *Witchcraft) Draw(screen *ebiten.Image) {
	screen.Clear()
	for _, s := range w.drawables {
		s.Draw(w.world, screen)
	}
}
