package component

import (
	"fmt"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

type ListData struct {
	Items  []*donburi.Entry
	Offset math.Vec2
}

func (d *ListData) Append(list *donburi.Entry, item *donburi.Entry) {
	offset := math.Vec2{
		X: d.Offset.X * float64(len(d.Items)),
		Y: d.Offset.Y * float64(len(d.Items)),
	}

	d.Items = append(d.Items, item)

	transform.GetTransform(item).LocalPosition = offset
	transform.AppendChild(list, item, false)
}

func (d *ListData) Remove(index int) {
	if index < 0 || index >= len(d.Items) {
		panic(fmt.Sprintf("index out of range: %d", index))
	}

	Destroy(d.Items[index])

	d.Items = append(d.Items[:index], d.Items[index+1:]...)

	for i, item := range d.Items {
		transform.GetTransform(item).LocalPosition = math.Vec2{
			X: d.Offset.X * float64(i),
			Y: d.Offset.Y * float64(i),
		}
	}
}

var List = donburi.NewComponentType[ListData]()
