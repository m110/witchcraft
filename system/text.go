package system

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
	"golang.org/x/image/colornames"
	stdfont "golang.org/x/image/font"

	"github.com/m110/witchcraft/assets"
	"github.com/m110/witchcraft/component"
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

func (t *Text) Update(w donburi.World) {
	t.query.Each(w, func(entry *donburi.Entry) {
		txt := component.Text.Get(entry)
		if txt.Update != nil {
			txt.Update(txt)
		}
	})
}

func (t *Text) Draw(w donburi.World, screen *ebiten.Image) {
	t.query.Each(w, func(entry *donburi.Entry) {
		t := component.Text.Get(entry)

		var font stdfont.Face
		switch t.Size {
		case component.TextSizeSmall:
			font = assets.SmallFont
		case component.TextSizeLarge:
			font = assets.NormalFont
		}

		pos := transform.WorldPosition(entry)

		var col color.Color = colornames.White
		if t.Color != nil {
			col = t.Color
		}

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(pos.X, pos.Y)
		op.ColorScale.ScaleWithColor(col)
		text.DrawWithOptions(screen, t.Text, font, op)
	})
}
