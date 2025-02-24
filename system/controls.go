package system

import (
	stdmath "math"

	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/hajimehoshi/ebiten/v2"
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
				component.Mover,
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

		moveSpeed := component.Mover.Get(entry).Speed

		delta := math.Vec2{}

		if input.GamepadID != nil {
			lx, ly := ebiten.GamepadAxisValue(*input.GamepadID, 0), ebiten.GamepadAxisValue(*input.GamepadID, 1)
			if stdmath.Abs(lx) > 0.2 || stdmath.Abs(ly) > 0.2 {
				delta.X += lx * moveSpeed
				delta.Y += ly * moveSpeed
			}

			rx, ry := ebiten.GamepadAxisValue(*input.GamepadID, 2), ebiten.GamepadAxisValue(*input.GamepadID, 5)
			if stdmath.Abs(rx) > 0.2 || stdmath.Abs(ry) > 0.2 {
				component.Direction.Get(entry).Direction = math.Vec2{X: rx, Y: ry}
			}
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

			// Handle mouse aiming
			cx, cy := ebiten.CursorPosition()
			t := transform.Transform.Get(entry)
			// TODO World position?
			dx := float64(cx) - t.LocalPosition.X
			dy := float64(cy) - t.LocalPosition.Y
			length := stdmath.Sqrt(dx*dx + dy*dy)
			if length > 0 {
				component.Direction.Get(entry).Direction = math.Vec2{
					X: dx / length,
					Y: dy / length,
				}
			}
		}

		// Check for diagonal movement
		if delta.X != 0 && delta.Y != 0 {
			factor := moveSpeed / stdmath.Sqrt(delta.X*delta.X+delta.Y*delta.Y)
			delta.X *= factor
			delta.Y *= factor
		}

		velocity := component.Velocity.Get(entry)
		velocity.Velocity = delta

		if entry.HasComponent(component.Caster) {
			caster := component.Caster.Get(entry)
			caster.IsCasting = ebiten.IsKeyPressed(input.CastKey) || input.GamepadID != nil && ebiten.IsStandardGamepadButtonPressed(*input.GamepadID, ebiten.StandardGamepadButtonFrontBottomRight)

			// TODO Needs a better logic when changing spells is allowed
			if inpututil.IsKeyJustPressed(input.SpellKeyA) || input.GamepadID != nil && ebiten.IsStandardGamepadButtonPressed(*input.GamepadID, ebiten.StandardGamepadButtonRightLeft) {
				caster.PrepareSpell(0)
			} else if inpututil.IsKeyJustPressed(input.SpellKeyB) || input.GamepadID != nil && ebiten.IsStandardGamepadButtonPressed(*input.GamepadID, ebiten.StandardGamepadButtonRightTop) {
				caster.PrepareSpell(1)
			} else if inpututil.IsKeyJustPressed(input.SpellKeyC) || input.GamepadID != nil && ebiten.IsStandardGamepadButtonPressed(*input.GamepadID, ebiten.StandardGamepadButtonRightRight) {
				caster.PrepareSpell(2)
			}
		}
	})
}
