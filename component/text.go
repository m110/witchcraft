package component

import (
	"image/color"

	"github.com/yohamta/donburi"
)

type TextSize int

const (
	TextSizeSmall TextSize = iota
	TextSizeLarge          = 1
)

type TextData struct {
	Size  TextSize
	Text  string
	Color color.Color
}

var Text = donburi.NewComponentType[TextData]()
