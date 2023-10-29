package system

import (
	"github.com/m110/witchcraft/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type TimeToLive struct {
	query *query.Query
}

func NewTimeToLive() *TimeToLive {
	return &TimeToLive{
		query: query.NewQuery(
			filter.Contains(component.TimeToLive),
		),
	}
}

func (t *TimeToLive) Update(w donburi.World) {
	t.query.Each(w, func(entry *donburi.Entry) {
		ttl := component.TimeToLive.Get(entry)
		ttl.Timer.Update()
		if ttl.Timer.IsReady() {
			w.Remove(entry.Entity())
		}
	})
}
