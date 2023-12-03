package system

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"

	"github.com/m110/witchcraft/component"
)

type Seeker struct {
	query        *donburi.Query
	playersQuery *donburi.Query
}

func NewSeeker() *Seeker {
	return &Seeker{
		query: donburi.NewQuery(
			filter.Contains(
				component.Seeker,
				component.Targeter,
			),
		),
		playersQuery: donburi.NewQuery(
			filter.Contains(
				component.Player,
			),
		),
	}
}

func (s *Seeker) Update(w donburi.World) {
	s.query.Each(w, func(entry *donburi.Entry) {
		seeker := component.Seeker.Get(entry)

		switch seeker.Type {
		case component.SeekerTypeNone:
		case component.SeekerTypeNearestPlayer:
			target := component.Targeter.Get(entry)

			target.Target = s.findNearestPlayer(w, entry)
		}
	})
}

func (s *Seeker) findNearestPlayer(w donburi.World, seeker *donburi.Entry) *donburi.Entry {
	var nearestPlayer *donburi.Entry
	var nearestDistance float64

	s.playersQuery.Each(w, func(entry *donburi.Entry) {
		playerPos := transform.WorldPosition(entry)
		seekerPos := transform.WorldPosition(seeker)

		distance := playerPos.Distance(seekerPos)
		if nearestPlayer == nil || distance < nearestDistance {
			nearestPlayer = entry
			nearestDistance = distance
		}
	})

	return nearestPlayer
}
