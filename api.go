package ebicanvas

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewCanvas(width, height int, stretched, pixelPerfect bool) *Canvas {
	result := &Canvas{
		width:        width,
		height:       height,
		stretched:    stretched,
		pixelPerfect: pixelPerfect,
		buffer:       ebiten.NewImage(width, height),
		renderers:    make([]func(*ebiten.Image), 0),
	}

	result.SetFilter(AASamplingSoft)

	return result
}

// GetHeight - return 16:9 height from specified width
func GetHeight(width int) int {
	return int(math.Floor(float64(width) / (16.0 / 9.0)))
}
