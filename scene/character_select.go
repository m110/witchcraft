package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"

	"github.com/m110/cliche-rpg/archetype"
	"github.com/m110/cliche-rpg/component"
	"github.com/m110/cliche-rpg/system"
)

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
	g.spawnCharacters()
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

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		query.NewQuery(filter.Contains(component.Character)).EachEntity(g.world, func(entry *donburi.Entry) {
			entry.Remove()
		})
		g.spawnCharacters()
	}
}

func (g *CharacterSelect) spawnCharacters() {
	offset := 48.0

	for i := 0; i < 15; i++ {
		for j := 0; j < 12; j++ {
			archetype.NewRandomCharacter(g.world, math.Vec2{X: offset + float64(i)*offset, Y: offset + float64(j)*offset})
		}
	}
}

func (g *CharacterSelect) Draw(screen *ebiten.Image) {
	screen.Clear()
	for _, s := range g.drawables {
		s.Draw(g.world, screen)
	}
}
