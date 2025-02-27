package system

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"

	"github.com/m110/witchcraft/component"
	"github.com/m110/witchcraft/spell"
)

type Velocity struct {
	query *query.Query
}

func NewVelocity() *Velocity {
	return &Velocity{
		query: query.NewQuery(
			filter.Contains(transform.Transform, component.Velocity),
		),
	}
}

func (v *Velocity) Update(w donburi.World) {
	v.query.Each(w, func(entry *donburi.Entry) {
		t := transform.Transform.Get(entry)
		velocity := component.Velocity.Get(entry)

		if entry.HasComponent(component.AuraHolder) {
			holder := component.AuraHolder.Get(entry)
			for _, aura := range holder.Auras {
				if aura.Effect.OnApply == spell.AuraEffectTypeSlowMovement {
					amount := aura.Effect.Amount
					velocity.Velocity = velocity.Velocity.MulScalar(amount)
				} else if aura.Effect.OnApply == spell.AuraEffectTypeSpeedBoost {
					amount := aura.Effect.Amount
					velocity.Velocity = velocity.Velocity.MulScalar(amount)
				}
			}
		}

		t.LocalPosition = t.LocalPosition.Add(velocity.Velocity)
		t.LocalRotation += velocity.RotationVelocity
	})
}
