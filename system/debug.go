package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"

	"github.com/m110/witchcraft/component"
)

type Debug struct {
	debug *component.DebugData
}

func NewDebug() *Debug {
	return &Debug{}
}

func (d *Debug) Update(w donburi.World) {
	if d.debug == nil {
		debug, ok := query.NewQuery(filter.Contains(component.Debug)).First(w)
		if !ok {
			return
		}
		d.debug = component.Debug.Get(debug)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySlash) {
		d.debug.Enabled = !d.debug.Enabled
	}
}
