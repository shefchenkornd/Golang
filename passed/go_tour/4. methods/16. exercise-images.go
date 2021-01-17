package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{}

func (i Image) ColorModel() color.Model  {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 30, 30)
}

func (i Image) At(x, y int) color.Color  {
	return color.RGBA{
		R: 117,
		G: 10,
		B: 10,
		A: 255,
	}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
