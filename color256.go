package rgbto256colors

import (
	"image/color"
	"math"

	"github.com/mattn/go-ciede2000"
)

type Color256 struct {
	ColorID   uint8
	Name      string
	HexString string
	R, G, B   uint8
}

func (c *Color256) ToColor() color.Color {
	return &color.RGBA{
		R: c.R,
		G: c.G,
		B: c.B,
		A: 255,
	}
}

func FromRGB(r, g, b uint8) *Color256 {
	t := &color.RGBA{r, g, b, 255}
	minDiff := math.MaxFloat64
	var bestColor *Color256
	for _, c := range colors {
		d := ciede2000.Diff(c.ToColor(), t)
		if d < minDiff {
			minDiff = d
			bestColor = c
		}
	}
	return bestColor
}
