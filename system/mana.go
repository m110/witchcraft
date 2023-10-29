package system

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"

	"github.com/m110/cliche-rpg/component"
)

type Mana struct {
	query *query.Query
}

func NewMana() *Mana {
	return &Mana{
		query: query.NewQuery(
			filter.Contains(component.Mana),
		),
	}
}

func (m *Mana) Update(w donburi.World) {
	m.query.EachEntity(w, func(entry *donburi.Entry) {
		mana := component.Mana.Get(entry)

		mana.ManaRegenTimer.Update()

		if mana.ManaRegenTimer.IsReady() {
			mana.Mana += mana.ManaRegen
			if mana.Mana > mana.MaxMana {
				mana.Mana = mana.MaxMana
			}

			mana.ManaRegenTimer.Reset()
		}
	})
}
