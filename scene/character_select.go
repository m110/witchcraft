package scene

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"

	"github.com/m110/witchcraft/archetype"
	"github.com/m110/witchcraft/component"
	"github.com/m110/witchcraft/system"
)

type CharacterSelect struct {
	context Context

	world     donburi.World
	systems   []System
	drawables []Drawable

	classes []archetype.Class
	players []PlayerSelect
}

func NewCharacterSelect(context Context) *CharacterSelect {
	g := &CharacterSelect{
		context: context,

		players: make([]PlayerSelect, 4),
	}

	g.loadLevel()

	return g
}

type PlayerSelect struct {
	Joined bool
	Ready  bool

	Keyboard  bool
	GamePadID *ebiten.GamepadID

	ClassIndex int
	Class      archetype.Class
	Demo       *donburi.Entry
}

func (p *PlayerSelect) Show(keyboard bool, gamePadID *ebiten.GamepadID, classIndex int, class archetype.Class) {
	p.Joined = true
	p.Keyboard = keyboard
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
	p.GamePadID = nil
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
			ScreenWidth:  g.context.ScreenWidth,
			ScreenHeight: g.context.ScreenHeight,
		},
	})

	world.Create(component.Debug)

	archetype.NewText(world, "Press (X) to join", component.TextSizeLarge, math.Vec2{X: float64(g.context.ScreenWidth) / 3.0, Y: 50})

	g.classes = archetype.LoadClasses()

	offsetX := float64(g.context.ScreenWidth) / 3.0
	offsetY := float64(g.context.ScreenHeight) / 3.0

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

		classNameText := archetype.NewText(world, "", component.TextSizeLarge, math.Vec2{X: -15, Y: 80})
		classNameText.AddComponent(component.ClassName)
		transform.AppendChild(demo, classNameText, false)

		readyText := archetype.NewText(world, "", component.TextSizeLarge, math.Vec2{X: -30, Y: -20})
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
		id := id
		if inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonRightBottom) {
			g.onPlayerJoin(false, &id)
		} else if inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonRightRight) {
			g.onPlayerLeave(false, &id)
		} else if inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonLeftLeft) {
			g.onChangeClass(false, &id, -1)
		} else if inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonLeftRight) {
			g.onChangeClass(false, &id, 1)
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.onPlayerJoin(true, nil)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.onPlayerLeave(true, nil)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		g.onChangeClass(true, nil, -1)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		g.onChangeClass(true, nil, 1)
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
				Keyboard:  p.Keyboard,
				GamePadID: p.GamePadID,
				Class:     p.Class,
			})
		}
	}

	g.context.SwitchToBattle(joinedPlayers)
}

func (g *CharacterSelect) onPlayerJoin(keyboard bool, gamepadID *ebiten.GamepadID) {
	for i, p := range g.players {
		if (keyboard && p.Keyboard) || (p.GamePadID != nil && *p.GamePadID == *gamepadID) {
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

	if keyboard {
		fmt.Printf("Player %v joined with keyboard\n", freeIndex)
	} else {
		fmt.Printf("Player %v joined with game pad ID %v\n", freeIndex, *gamepadID)
	}

	class := g.classes[0]

	g.players[freeIndex].Show(keyboard, gamepadID, 0, class)
}

func (g *CharacterSelect) onPlayerLeave(keyboard bool, gamepadID *ebiten.GamepadID) {
	var anyJoined bool
	for _, p := range g.players {
		if p.Joined {
			anyJoined = true
			break
		}
	}

	if !anyJoined {
		g.context.SwitchToMainMenu()
		return
	}

	for i, p := range g.players {
		if p.Joined &&
			((keyboard && p.Keyboard) || (p.GamePadID != nil && *p.GamePadID == *gamepadID)) {
			if p.Ready {
				g.players[i].Unready()
			} else {
				g.players[i].Hide()

				if keyboard {
					fmt.Printf("Player %v left with keyboard\n", i)
				} else {
					fmt.Printf("Player %v left with game pad ID %v\n", i, *gamepadID)
				}
			}
		}
	}
}

func (g *CharacterSelect) onChangeClass(keyboard bool, gamepadID *ebiten.GamepadID, direction int) {
	for i, p := range g.players {
		if p.Joined && !p.Ready &&
			((keyboard && p.Keyboard) || (p.GamePadID != nil && *p.GamePadID == *gamepadID)) {
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
