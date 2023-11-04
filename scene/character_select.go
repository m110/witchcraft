package scene

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/m110/witchcraft/archetype"
	"github.com/m110/witchcraft/component"
	"github.com/m110/witchcraft/system"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

type CharacterSelect struct {
	world     donburi.World
	systems   []System
	drawables []Drawable

	classes []archetype.Class
	players []Player

	screenWidth  int
	screenHeight int
}

func NewCharacterSelect(screenWidth int, screenHeight int) *CharacterSelect {
	g := &CharacterSelect{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,

		players: make([]Player, 4),
	}

	g.loadLevel()

	return g
}

type Player struct {
	Joined    bool
	GamePadID ebiten.GamepadID
	Class     archetype.Class
	Demo      *donburi.Entry
}

func (p *Player) Show(gamePadID ebiten.GamepadID, class archetype.Class) {
	p.Joined = true
	p.GamePadID = gamePadID
	p.SetClass(class)
}

func (p *Player) SetClass(class archetype.Class) {
	p.Class = class

	sprite := component.Sprite.Get(p.Demo)
	sprite.Hidden = false
	sprite.Image = class.Character.Image()

	child, ok := transform.FindChildWithComponent(p.Demo, component.Text)
	if ok {
		component.Text.Get(child).Text = class.Name
	}
}

func (p *Player) Hide() {
	p.Joined = false
	p.GamePadID = 0
	p.Class = archetype.Class{}

	component.Sprite.Get(p.Demo).Hidden = true

	child, ok := transform.FindChildWithComponent(p.Demo, component.Text)
	if ok {
		component.Text.Get(child).Text = ""
	}
}

func (g *CharacterSelect) loadLevel() {
	render := system.NewRenderer()

	g.systems = []System{
		render,
	}

	g.drawables = []Drawable{
		system.NewText(),
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

	archetype.NewText(world, "Press (X) to join", component.TextSizeLarge, math.Vec2{X: float64(g.screenWidth) / 3.0, Y: 50})

	g.classes = archetype.LoadClasses()

	offsetX := float64(g.screenWidth) / 3.0
	offsetY := float64(g.screenHeight) / 3.0

	playerPositions := []math.Vec2{
		{X: offsetX, Y: offsetY},
		{X: offsetX * 2, Y: offsetY},
		{X: offsetX, Y: offsetY * 2},
		{X: offsetX * 2, Y: offsetY * 2},
	}

	for i := range g.players {
		demo := world.Entry(
			world.Create(
				transform.Transform,
				component.Sprite,
			),
		)
		transform.GetTransform(demo).LocalPosition = playerPositions[i]
		transform.GetTransform(demo).LocalScale = math.Vec2{X: 3, Y: 3}
		component.Sprite.Get(demo).Hidden = true
		text := archetype.NewText(world, "", component.TextSizeLarge, math.Vec2{X: -30, Y: 60})
		transform.AppendChild(demo, text, false)

		g.players[i] = Player{
			Demo: demo,
		}
	}

	return world
}

func (g *CharacterSelect) Update() {
	for _, id := range ebiten.AppendGamepadIDs(nil) {
		if ebiten.IsStandardGamepadButtonPressed(id, ebiten.StandardGamepadButtonRightBottom) {
			g.onPlayerJoin(id)
		} else if ebiten.IsStandardGamepadButtonPressed(id, ebiten.StandardGamepadButtonRightRight) {
			g.onPlayerLeave(id)
		}
	}

	for _, s := range g.systems {
		s.Update(g.world)
	}
}

func (g *CharacterSelect) onPlayerJoin(id ebiten.GamepadID) {
	for _, p := range g.players {
		if p.Joined && p.GamePadID == id {
			// Already joined
			return
		}
	}

	free := false
	freeIndex := 0
	for i, p := range g.players {
		if !p.Joined {
			free = true
			freeIndex = i
			fmt.Println("Free player slot", i)
			break
		}
	}

	if !free {
		return
	}

	fmt.Printf("Player %v joined with game pad ID %v\n", freeIndex, id)

	class := g.classes[0]

	g.players[freeIndex].Show(id, class)
}

func (g *CharacterSelect) onPlayerLeave(id ebiten.GamepadID) {
	for i, p := range g.players {
		if p.Joined && p.GamePadID == id {
			g.players[i].Hide()
			fmt.Printf("Player %v left with game pad ID %v\n", i, id)
		}
	}
}

func (g *CharacterSelect) Draw(screen *ebiten.Image) {
	for _, s := range g.drawables {
		s.Draw(g.world, screen)
	}
}
