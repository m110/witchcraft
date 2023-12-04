package system

import (
	"math"
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/samber/lo"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
	"golang.org/x/image/colornames"

	"github.com/m110/witchcraft/archetype"
	"github.com/m110/witchcraft/component"
)

type Render struct {
	query     *query.Query
	offscreen *ebiten.Image
	debug     *component.DebugData
}

func NewRenderer() *Render {
	return &Render{
		query: query.NewQuery(
			filter.Contains(transform.Transform, component.Sprite),
		),
		// TODO figure out the proper size
		offscreen: ebiten.NewImage(3000, 3000),
	}
}

func (r *Render) Update(w donburi.World) {
	if r.debug == nil {
		debug, ok := query.NewQuery(filter.Contains(component.Debug)).First(w)
		if !ok {
			return
		}

		r.debug = component.Debug.Get(debug)
	}
}

func (r *Render) Draw(w donburi.World, screen *ebiten.Image) {
	camera := archetype.MustFindCamera(w)
	cameraPos := transform.GetTransform(camera).LocalPosition

	r.offscreen.Clear()

	var entries []*donburi.Entry
	r.query.Each(w, func(entry *donburi.Entry) {
		entries = append(entries, entry)
	})

	byLayer := lo.GroupBy(entries, func(entry *donburi.Entry) int {
		return int(component.Sprite.Get(entry).Layer)
	})
	layers := lo.Keys(byLayer)
	sort.Ints(layers)

	for _, layer := range layers {
		for _, entry := range byLayer[layer] {
			sprite := component.Sprite.Get(entry)

			if sprite.Hidden {
				continue
			}

			bounds := sprite.Image.Bounds()

			width, height := bounds.Dx(), bounds.Dy()

			halfW, halfH := float64(width)/2, float64(height)/2

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(-halfW, -halfH)
			op.GeoM.Rotate(float64(int(transform.WorldRotation(entry)-sprite.OriginalRotation)%360) * 2 * math.Pi / 360)
			op.GeoM.Translate(halfW, halfH)

			position := transform.WorldPosition(entry)

			x := position.X
			y := position.Y

			scale := transform.WorldScale(entry)

			switch sprite.Pivot {
			case component.SpritePivotCenter:
				x -= halfW * scale.X
				y -= halfH * scale.Y
			}

			op.GeoM.Scale(scale.X, scale.Y)

			if sprite.ColorOverride != nil {
				op.ColorM.Scale(0, 0, 0, sprite.ColorOverride.A)
				op.ColorM.Translate(sprite.ColorOverride.R, sprite.ColorOverride.G, sprite.ColorOverride.B, 0)
			}

			op.GeoM.Translate(x, y)

			r.offscreen.DrawImage(sprite.Image, op)

			if r.debug != nil && r.debug.Enabled {
				if entry.HasComponent(component.Collider) {
					collider := component.Collider.Get(entry)
					rect := collider.Rect(entry)
					vector.StrokeRect(
						r.offscreen,
						float32(rect.X), float32(rect.Y),
						float32(rect.Width), float32(rect.Height),
						1,
						colornames.Lime,
						true,
					)

					for collision := range collider.Collisions {
						other := w.Entry(collision.Other)
						if !other.Valid() {
							continue
						}
						otherPos := transform.WorldPosition(other)
						vector.StrokeLine(
							r.offscreen,
							float32(position.X), float32(position.Y),
							float32(otherPos.X), float32(otherPos.Y),
							1,
							colornames.Yellow,
							true,
						)
					}
				}
			}
		}
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-cameraPos.X, -cameraPos.Y)
	screen.DrawImage(r.offscreen, op)
}
