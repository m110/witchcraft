package scene

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/colornames"

	"github.com/m110/witchcraft/assets"
)

type MenuItem struct {
	Text   string
	Action func()
}

type MainMenu struct {
	context Context

	menuItems       []MenuItem
	activeItemIndex int
}

func NewMainMenu(context Context) *MainMenu {
	return &MainMenu{
		context: context,
		menuItems: []MenuItem{
			{
				Text: "Play",
				Action: func() {
					context.SwitchToCharacterSelect()
				},
			},
			{
				Text: "Fitting Room",
				Action: func() {
					context.SwitchToFittingRoom()
				},
			},
			{
				Text: "Quit",
				Action: func() {
					os.Exit(0)
				},
			},
		},
		activeItemIndex: 0,
	}
}

func (m *MainMenu) Update() {
	var action, moveDown, moveUp bool

	for _, id := range ebiten.AppendGamepadIDs(nil) {
		if inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonRightBottom) {
			action = true
			break
		}

		if inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonLeftBottom) {
			moveDown = true
			break
		}

		if inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonLeftTop) {
			moveUp = true
			break
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		action = true
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		moveDown = true
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		moveUp = true
	}

	if action {
		m.menuItems[m.activeItemIndex].Action()
	} else if moveDown {
		m.activeItemIndex++
		if m.activeItemIndex >= len(m.menuItems) {
			m.activeItemIndex = 0
		}
	} else if moveUp {
		m.activeItemIndex--
		if m.activeItemIndex < 0 {
			m.activeItemIndex = len(m.menuItems) - 1
		}
	}
}

func (m *MainMenu) Draw(screen *ebiten.Image) {
	for i, item := range m.menuItems {
		col := colornames.White
		if i == m.activeItemIndex {
			col = colornames.Yellow
		}
		text.Draw(screen, item.Text, assets.NormalFont, 100, int(100+float64(i)*50), col)
	}
}
