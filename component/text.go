package component

import "github.com/yohamta/donburi"

type TextSize int

const (
	TextSizeSmall TextSize = iota
	TextSizeLarge          = 1
)

type TextData struct {
	Size TextSize
	Text string
}

var Text = donburi.NewComponentType[TextData]()
