package component

import "github.com/yohamta/donburi"

type TextData struct {
	Text string
}

var Text = donburi.NewComponentType[TextData]()
