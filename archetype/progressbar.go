package archetype

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"

	"github.com/m110/cliche-rpg/component"
)

func NewProgressBar(
	w donburi.World,
	pos math.Vec2,
	width int,
	height int,
	backgroundColor color.Color,
	update func(bar *component.ProgressBarData),
) *donburi.Entry {
	bar := w.Entry(
		w.Create(
			transform.Transform,
			component.ProgressBar,
			component.Sprite,
		),
	)

	transform.GetTransform(bar).LocalPosition = pos

	component.ProgressBar.Set(bar, &component.ProgressBarData{
		Width:           width,
		Height:          height,
		BackgroundColor: backgroundColor,
		Update:          update,
	})

	component.Sprite.Set(bar, &component.SpriteData{
		Image: ebiten.NewImage(width, height),
		Pivot: component.SpritePivotTopLeft,
		Layer: component.SpriteLayerUI,
	})

	return bar
}
