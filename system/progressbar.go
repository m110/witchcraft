package system

import (
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"

	"github.com/m110/cliche-rpg/component"
)

type ProgressBar struct {
	query *query.Query
}

func NewProgressBar() *ProgressBar {
	return &ProgressBar{
		query: query.NewQuery(
			filter.Contains(
				transform.Transform,
				component.ProgressBar,
				component.Sprite,
			),
		),
	}
}

func (p *ProgressBar) Update(w donburi.World) {
	p.query.EachEntity(w, func(entry *donburi.Entry) {
		bar := component.ProgressBar.Get(entry)

		value := bar.Value
		maxValue := bar.MaxValue
		bgColor := bar.BackgroundColor

		bar.Update(bar)

		// Nothing changed
		if value == bar.Value && maxValue == bar.MaxValue && bgColor == bar.BackgroundColor {
			return
		}

		sprite := component.Sprite.Get(entry)
		sprite.Image.Clear()

		width := float64(bar.Width) * (float64(bar.Value) / float64(bar.MaxValue))

		ebitenutil.DrawRect(
			sprite.Image,
			0,
			0,
			width,
			float64(bar.Height),
			bar.BackgroundColor,
		)
	})
}
