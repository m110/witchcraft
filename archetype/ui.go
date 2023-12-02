package archetype

import (
	"image/color"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"

	"github.com/m110/witchcraft/component"
)

func NewText(w donburi.World, text string, size component.TextSize, pos math.Vec2) *donburi.Entry {
	t := w.Entry(
		w.Create(
			transform.Transform,
			component.Text,
		),
	)

	transform.GetTransform(t).LocalPosition = pos

	component.Text.Set(t, &component.TextData{
		Size: size,
		Text: text,
	})

	return t
}

func NewAuraIcon(w donburi.World, aura component.Aura) *donburi.Entry {
	icon := w.Entry(w.Create(
		transform.Transform,
		component.Sprite,
	))

	component.Sprite.SetValue(icon, component.SpriteData{
		Image: aura.Template.Image,
		Pivot: component.SpritePivotTopLeft,
		Layer: component.SpriteLayerUI,
	})

	pb := NewProgressBar(w, math.Vec2{X: 0, Y: 31}, 32, 6, color.White, func(bar *component.ProgressBarData) {
		bar.Value = 100 - int(aura.Timer.PercentDone()*100)
		bar.MaxValue = 100
	})
	transform.AppendChild(icon, pb, false)

	return icon
}
