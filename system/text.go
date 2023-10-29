package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/m110/witchcraft/assets"
	"github.com/m110/witchcraft/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
	"golang.org/x/image/colornames"
)

type Text struct {
	query *query.Query
}

func NewText() *Text {
	return &Text{
		query: query.NewQuery(
			filter.Contains(
				transform.Transform,
				component.Text,
			),
		),
	}
}

func (t *Text) Draw(w donburi.World, screen *ebiten.Image) {
	t.query.Each(w, func(entry *donburi.Entry) {
		transform := transform.Transform.Get(entry)
		t := component.Text.Get(entry)

		text.Draw(screen, t.Text, assets.SmallFont, int(transform.LocalPosition.X), int(transform.LocalPosition.Y), colornames.White)
	})
}
