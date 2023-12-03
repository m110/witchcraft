package archetype

import (
	"fmt"
	"image/color"
	"time"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
	"golang.org/x/image/colornames"

	"github.com/m110/witchcraft/component"
	"github.com/m110/witchcraft/engine"
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

	if aura.Timer != nil {
		pb := NewProgressBar(w, math.Vec2{X: 0, Y: 31}, 32, 6, color.White, func(bar *component.ProgressBarData) {
			bar.Value = 100 - int(aura.Timer.PercentDone()*100)
			bar.MaxValue = 100
		})
		transform.AppendChild(icon, pb, false)
	}

	return icon
}

func NewDamageText(w donburi.World, damage int, pos math.Vec2) *donburi.Entry {
	t := w.Entry(
		w.Create(
			transform.Transform,
			component.Text,
			component.TimeToLive,
			component.Velocity,
		),
	)

	transform.GetTransform(t).LocalPosition = pos

	component.Text.Set(t, &component.TextData{
		Size:  component.TextSizeSmall,
		Text:  fmt.Sprint(damage),
		Color: colornames.Red,
	})

	component.TimeToLive.Set(t, &component.TimeToLiveData{
		Timer: engine.NewTimer(time.Second),
	})

	component.Velocity.Set(t, &component.VelocityData{
		Velocity: math.Vec2{X: 0, Y: -1},
	})

	return t
}
