package rgbto256colors

import (
	"testing"
)

type rgb struct {
	r, g, b uint8
}

func TestFromRGB(t *testing.T) {
	existColors := []*rgb{
		&rgb{r: 0, g: 0, b: 255},    // 21	Blue1	#0000ff
		&rgb{r: 0, g: 255, b: 0},    // 46	Green1	#00ff00
		&rgb{r: 0, g: 255, b: 255},  // 51	Cyan1	#00ffff
		&rgb{r: 95, g: 255, b: 135}, // 84	SeaGreen1	#5fff87
		&rgb{r: 175, g: 0, b: 255},  // 129	Purple	#af00ff
		&rgb{r: 255, g: 0, b: 0},    // 196	Red1	#ff0000
		&rgb{r: 255, g: 215, b: 0},  // 220	Gold1	#ffd700
	}
	for _, c := range existColors {
		c256 := FromRGB(c.r, c.g, c.b)
		if c256.R != c.r || c256.G != c.g || c256.B != c.b {
			t.Fatalf("from: %v, created: %v", c, c256)
		}
	}
}
