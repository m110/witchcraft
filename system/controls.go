package system

import (
	stdmath "math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"

	"github.com/m110/witchcraft/component"
)

type Controls struct {
	query *query.Query
}

func NewControls() *Controls {
	return &Controls{
		query: query.NewQuery(
			filter.Contains(
				transform.Transform,
				component.Velocity,
				component.Input,
			),
		),
	}
}

func (i *Controls) Update(w donburi.World) {
	i.query.Each(w, func(entry *donburi.Entry) {
		input := component.Input.Get(entry)

		if input.Disabled {
			return
		}

		// TODO Move this somewhere to a component
		const moveSpeed = 3

		delta := math.Vec2{}

		lx, ly := ebiten.GamepadAxisValue(0, 0), ebiten.GamepadAxisValue(0, 1)
		if stdmath.Abs(lx) > 0.2 || stdmath.Abs(ly) > 0.2 {
			delta.X += lx * moveSpeed
			delta.Y += ly * moveSpeed
		} else {
			if ebiten.IsKeyPressed(input.MoveUpKey) {
				delta.Y = -moveSpeed
			} else if ebiten.IsKeyPressed(input.MoveDownKey) {
				delta.Y = moveSpeed
			}

			if ebiten.IsKeyPressed(input.MoveRightKey) {
				delta.X = moveSpeed
			}
			if ebiten.IsKeyPressed(input.MoveLeftKey) {
				delta.X = -moveSpeed
			}

			// Check for diagonal movement
			if delta.X != 0 && delta.Y != 0 {
				factor := moveSpeed / stdmath.Sqrt(delta.X*delta.X+delta.Y*delta.Y)
				delta.X *= factor
				delta.Y *= factor
			}
		}

		velocity := component.Velocity.Get(entry)
		velocity.Velocity = delta

		rx, ry := ebiten.GamepadAxisValue(0, 2), ebiten.GamepadAxisValue(0, 5)
		if stdmath.Abs(rx) > 0.2 || stdmath.Abs(ry) > 0.2 {
			component.Direction.Get(entry).Direction = math.Vec2{X: rx, Y: ry}
		}

		if entry.HasComponent(component.Caster) {
			caster := component.Caster.Get(entry)
			caster.IsCasting = ebiten.IsKeyPressed(input.CastKey) || ebiten.IsStandardGamepadButtonPressed(0, ebiten.StandardGamepadButtonRightBottom)

			// TODO Needs a better logic when changing spells is allowed
			if inpututil.IsKeyJustPressed(input.SpellKeyA) {
				caster.PrepareSpell(0)
			} else if inpututil.IsKeyJustPressed(input.SpellKeyB) {
				caster.PrepareSpell(1)
			} else if inpututil.IsKeyJustPressed(input.SpellKeyC) {
				caster.PrepareSpell(2)
			}
		}
	})
}
