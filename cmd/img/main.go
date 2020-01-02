package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
	"time"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"

	"github.com/lusingander/rgbto256colors-go"
)

const output = "./output.png"

const (
	cellW    = 80
	cellWBuf = 30
	cellH    = 50
	cellHBuf = 20

	marginX = 10
	marginY = 10

	row = 10
	col = 3

	width  = (marginX * (col - 1)) + (cellWBuf * (col - 1)) + ((cellW * 2) * col)
	height = marginY + (cellH+cellHBuf)*row
)

var (
	white = color.RGBA{255, 255, 255, 255}
)

func fill(i *image.RGBA, c color.Color) {
	rect := i.Rect
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			i.Set(x, y, c)
		}
	}
}

func paintPattern(i *image.RGBA, r, c int, c1, c2 color.Color) {
	baseY := marginY + (cellH+cellHBuf)*r
	baseX := marginX + (cellW*2+cellWBuf)*c
	for y := baseY; y < baseY+cellH; y++ {
		for x := baseX; x < baseX+cellW; x++ {
			i.Set(x, y, c1)
		}
		for x := baseX + cellW; x < baseX+(cellW*2); x++ {
			i.Set(x, y, c2)
		}
	}
	drawColorCodeString(i, baseX, baseY, c1, c2)
}

func drawColorCodeString(i *image.RGBA, baseX, baseY int, c1, c2 color.Color) {
	d := &font.Drawer{
		Dst:  i,
		Src:  image.NewUniform(color.Black),
		Face: basicfont.Face7x13,
	}
	d.Dot.X = (fixed.I(baseX))
	d.Dot.Y = fixed.I(baseY + cellH + d.Face.Metrics().Height.Floor())
	d.DrawString(hexColorString(c1))
	d.Dot.X = (fixed.I(baseX + cellW))
	d.Dot.Y = fixed.I(baseY + cellH + d.Face.Metrics().Height.Floor())
	d.DrawString(hexColorString(c2))
}

func hexColorString(c color.Color) string {
	rgba := color.RGBAModel.Convert(c).(color.RGBA)
	return fmt.Sprintf("#%.2X%.2X%.2X", rgba.R, rgba.G, rgba.B)
}

func uint8RandGen() func() uint8 {
	rand.Seed(time.Now().UnixNano())
	return func() uint8 { return uint8(rand.Intn(255)) }
}

func draw(img *image.RGBA) {
	fill(img, white)

	var randN = uint8RandGen()

	for c := 0; c < col; c++ {
		for r := 0; r < row; r++ {
			c1 := color.RGBA{randN(), randN(), randN(), 255}
			c2 := rgbto256colors.FromRGB(c1.R, c1.G, c1.B).ToColor()
			paintPattern(img, r, c, c1, c2)
		}
	}
}

func save(img *image.RGBA) error {
	file, err := os.Create(output)
	if err != nil {
		return err
	}
	defer file.Close()
	err = png.Encode(file, img)
	if err != nil {
		return err
	}
	return nil
}

func run(args []string) error {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw(img)
	return save(img)
}

func main() {
	if err := run(os.Args); err != nil {
		log.Fatal(err)
	}
}
