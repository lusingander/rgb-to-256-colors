package rgbto256colors

import "image/color"

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
