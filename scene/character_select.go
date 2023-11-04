package scene

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	players []PlayerSelect

	screenWidth  int
	screenHeight int

	startBattleFunc func([]JoinedPlayer)
}

func NewCharacterSelect(screenWidth int, screenHeight int, startBattleFunc func([]JoinedPlayer)) *CharacterSelect {
	g := &CharacterSelect{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,

		players: make([]PlayerSelect, 4),

		startBattleFunc: startBattleFunc,
	}

	g.loadLevel()

	return g
}

type PlayerSelect struct {
	Joined     bool
	Ready      bool
	GamePadID  ebiten.GamepadID
	ClassIndex int
	Class      archetype.Class
	Demo       *donburi.Entry
}

func (p *PlayerSelect) Show(gamePadID ebiten.GamepadID, classIndex int, class archetype.Class) {
	p.Joined = true
	p.GamePadID = gamePadID
	p.SetClass(classIndex, class)
}

func (p *PlayerSelect) SetClass(classIndex int, class archetype.Class) {
	p.ClassIndex = classIndex
	p.Class = class

	sprite := component.Sprite.Get(p.Demo)
	sprite.Hidden = false
	sprite.Image = class.Character.Image()

	child, ok := transform.FindChildWithComponent(p.Demo, component.Text)
	if ok {
		component.Text.Get(child).Text = class.Name
	}
}

func (p *PlayerSelect) Hide() {
	p.Joined = false
	p.Ready = false
	p.GamePadID = 0
	p.Class = archetype.Class{}

	component.Sprite.Get(p.Demo).Hidden = true

	className, ok := transform.FindChildWithComponent(p.Demo, component.ClassName)
	if ok {
		component.Text.Get(className).Text = ""
	}
}

func (p *PlayerSelect) ReadyUp() {
	p.Ready = true

	ready, ok := transform.FindChildWithComponent(p.Demo, component.ReadyIndicator)
	if ok {
		component.Text.Get(ready).Text = "Ready"
	}
}

func (p *PlayerSelect) Unready() {
	p.Ready = false

	ready, ok := transform.FindChildWithComponent(p.Demo, component.ReadyIndicator)
	if ok {
		component.Text.Get(ready).Text = ""
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

		classNameText := archetype.NewText(world, "", component.TextSizeLarge, math.Vec2{X: -30, Y: 60})
		classNameText.AddComponent(component.ClassName)
		transform.AppendChild(demo, classNameText, false)

		readyText := archetype.NewText(world, "", component.TextSizeLarge, math.Vec2{X: -50, Y: -60})
		readyText.AddComponent(component.ReadyIndicator)
		transform.AppendChild(demo, readyText, false)

		g.players[i] = PlayerSelect{
			Demo: demo,
		}
	}

	return world
}

func (g *CharacterSelect) Update() {
	for _, id := range ebiten.AppendGamepadIDs(nil) {
		if inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonRightBottom) {
			g.onPlayerJoin(id)
		} else if inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonRightRight) {
			g.onPlayerLeave(id)
		} else if inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonLeftLeft) {
			g.onChangeClass(id, -1)
		} else if inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonLeftRight) {
			g.onChangeClass(id, 1)
		}
	}

	for _, s := range g.systems {
		s.Update(g.world)
	}

	g.checkAllPlayersReady()
}

func (g *CharacterSelect) checkAllPlayersReady() {
	joined := 0
	ready := 0
	for _, p := range g.players {
		if p.Joined {
			joined++

			if p.Ready {
				ready++
			}
		}
	}

	if joined == 0 || joined != ready {
		return
	}

	var joinedPlayers []JoinedPlayer
	for _, p := range g.players {
		if p.Joined {
			joinedPlayers = append(joinedPlayers, JoinedPlayer{
				GamePadID: p.GamePadID,
				Class:     p.Class,
			})
		}
	}

	g.startBattleFunc(joinedPlayers)
}

func (g *CharacterSelect) onPlayerJoin(id ebiten.GamepadID) {
	for i, p := range g.players {
		if p.Joined && p.GamePadID == id {
			g.players[i].ReadyUp()
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

	g.players[freeIndex].Show(id, 0, class)
}

func (g *CharacterSelect) onPlayerLeave(id ebiten.GamepadID) {
	for i, p := range g.players {
		if p.Joined && p.GamePadID == id {
			if p.Ready {
				g.players[i].Unready()
			} else {
				g.players[i].Hide()
				fmt.Printf("Player %v left with game pad ID %v\n", i, id)
			}
		}
	}
}

func (g *CharacterSelect) onChangeClass(id ebiten.GamepadID, direction int) {
	for i, p := range g.players {
		if p.Joined && p.GamePadID == id && !p.Ready {
			classIndex := p.ClassIndex + direction
			if classIndex < 0 {
				classIndex = len(g.classes) - 1
			} else if classIndex >= len(g.classes) {
				classIndex = 0
			}

			g.players[i].SetClass(classIndex, g.classes[classIndex])
		}
	}
}

func (g *CharacterSelect) Draw(screen *ebiten.Image) {
	for _, s := range g.drawables {
		s.Draw(g.world, screen)
	}
}
