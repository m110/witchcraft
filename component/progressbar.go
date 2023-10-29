package component

import (
	"image/color"

	"github.com/yohamta/donburi"
)

type ProgressBarData struct {
	MaxValue int
	Value    int

	Width  int
	Height int

	BackgroundColor color.Color

	Update func(bar *ProgressBarData)
}

var ProgressBar = donburi.NewComponentType[ProgressBarData]()
