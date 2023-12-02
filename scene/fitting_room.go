package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type FittingRoom struct {
	context Context
}

func NewFittingRoom(context Context) *FittingRoom {
	return &FittingRoom{
		context: context,
	}
}

func (f *FittingRoom) Update() {
	for _, id := range ebiten.AppendGamepadIDs(nil) {
		if inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonRightRight) {
			f.context.SwitchToMainMenu()
		}
	}
}

func (f *FittingRoom) Draw(screen *ebiten.Image) {

}
