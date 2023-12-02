package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type Context struct {
	ScreenWidth  int
	ScreenHeight int

	SwitchToMainMenu        func()
	SwitchToFittingRoom     func()
	SwitchToCharacterSelect func()
	SwitchToBattle          func([]JoinedPlayer)
}

type System interface {
	Update(w donburi.World)
}

type Drawable interface {
	Draw(w donburi.World, screen *ebiten.Image)
}
