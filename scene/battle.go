package scene

import (
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

	g.loadLevel()

	return g
}

func (b *Battle) loadLevel() {
	render := system.NewRenderer()

	b.systems = []System{
		system.NewVelocity(),
		system.NewControls(),
		system.NewCasting(),
		system.NewProgressBar(),
		system.NewMana(),
		system.NewCrosshair(),
		render,
		system.NewTimeToLive(),
	}

	b.drawables = []Drawable{
		render,
		system.NewText(),
	}

	b.world = b.createWorld()
	b.spawnCharacters()
}

func (b *Battle) createWorld() donburi.World {
	world := donburi.NewWorld()

	archetype.NewCamera(world, math.Vec2{})

	game := world.Entry(world.Create(component.Game))
	donburi.SetValue(game, component.Game, component.GameData{
		Settings: component.Settings{
			ScreenWidth:  b.screenWidth,
			ScreenHeight: b.screenHeight,
		},
	})

	world.Create(component.Debug)

	return world
}

func (b *Battle) Update() {
	for _, s := range b.systems {
		s.Update(b.world)
	}
}

func (b *Battle) spawnCharacters() {
	offset := 150.0
	positions := []math.Vec2{
		{X: offset, Y: offset},
		{X: float64(b.screenWidth) - offset, Y: offset},
		{X: offset, Y: float64(b.screenHeight) - offset},
		{X: float64(b.screenWidth) - offset, Y: float64(b.screenHeight) - offset},
	}

	for i, p := range b.joinedPlayers {
		archetype.NewPlayer(b.world, i, p.GamePadID, positions[i], p.Class)
	}
}

func (b *Battle) Draw(screen *ebiten.Image) {
	for _, s := range b.drawables {
		s.Draw(b.world, screen)
	}
}
