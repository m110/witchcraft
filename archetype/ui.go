package archetype

import (
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
