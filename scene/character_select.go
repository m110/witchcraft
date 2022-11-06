package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"

	"github.com/m110/cliche-rpg/archetype"
	"github.com/m110/cliche-rpg/component"
	"github.com/m110/cliche-rpg/system"
)

type System interface {
	Update(w donburi.World)
}

type Drawable interface {
	Draw(w donburi.World, screen *ebiten.Image)
}

type CharacterSelect struct {
	world     donburi.World
	systems   []System
	drawables []Drawable

	screenWidth  int
	screenHeight int
}

func NewCharacterSelect(screenWidth int, screenHeight int) *CharacterSelect {
	g := &CharacterSelect{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
	}

	g.loadLevel()

	return g
}

func (g *CharacterSelect) loadLevel() {
	render := system.NewRenderer()

	g.systems = []System{
		render,
	}

	g.drawables = []Drawable{
		render,
	}

	g.world = g.createWorld()
}

func (g *CharacterSelect) createWorld() donburi.World {
	world := donburi.NewWorld()

	archetype.NewCamera(world, math.Vec2{})

	game := world.Entry(world.Create(component.Game))
	donburi.SetValue(game, component.Game, component.GameData{
		Settings: component.Settings{
			ScreenWidth:  g.screenWidth,
			ScreenHeight: g.screenHeight,
		},
	})

	world.Create(component.Debug)

	return world
}

func (g *CharacterSelect) Update() {
	for _, s := range g.systems {
		s.Update(g.world)
	}
}

func (g *CharacterSelect) Draw(screen *ebiten.Image) {
	screen.Clear()
	for _, s := range g.drawables {
		s.Draw(g.world, screen)
	}
}
