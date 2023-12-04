package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"

	"github.com/m110/witchcraft/archetype"
	"github.com/m110/witchcraft/component"
	"github.com/m110/witchcraft/system"
)

type JoinedPlayer struct {
	GamePadID ebiten.GamepadID
	Class     archetype.Class
}

type Battle struct {
	context Context

	world     donburi.World
	systems   []System
	drawables []Drawable

	joinedPlayers []JoinedPlayer
}

func NewBattle(context Context, joinedPlayers []JoinedPlayer) *Battle {
	g := &Battle{
		context: context,

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
		system.NewAuras(),
		system.NewProgressBar(),
		system.NewMana(),
		system.NewCrosshair(),
		system.NewSpawner(),
		system.NewSeeker(),
		system.NewAI(),
		system.NewTimeToLive(),
		// Order matters: collisions expect entities to be marked as Destroyed
		// TODO: Probably not ready for the "out of collision due to collision damage" scenario
		system.NewCollision(),
		system.NewCollisionDamage(),
		system.NewCollisionApplyAura(),
		render,
		system.NewDestroy(),
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
			ScreenWidth:  b.context.ScreenWidth,
			ScreenHeight: b.context.ScreenHeight,
		},
	})

	s := archetype.NewSpawner(world)
	transform.GetTransform(s).LocalPosition = math.Vec2{X: 400, Y: 300}

	world.Create(component.Debug)

	return world
}

func (b *Battle) Update() {
	for _, p := range b.joinedPlayers {
		if inpututil.IsStandardGamepadButtonJustPressed(p.GamePadID, ebiten.StandardGamepadButtonCenterRight) {
			b.context.SwitchToCharacterSelect()
		}
	}

	for _, s := range b.systems {
		s.Update(b.world)
	}
}

func (b *Battle) spawnCharacters() {
	offset := 150.0
	positions := []math.Vec2{
		{X: offset, Y: offset},
		{X: float64(b.context.ScreenWidth) - offset, Y: offset},
		{X: offset, Y: float64(b.context.ScreenHeight) - offset},
		{X: float64(b.context.ScreenWidth) - offset, Y: float64(b.context.ScreenHeight) - offset},
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
