package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

const (
	width  int = 256
	height int = 256
)

type Image struct {
	w, h int
	v    []uint8
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i Image) At(x, y int) color.Color {
	v := i.v[width*y+x]
	return color.RGBA{v, 0, v, 255}
}

func main() {
	var data [width * height]uint8

	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			//fmt.Println(w ^ h)
			data[width*h+w] = uint8(w ^ h)
		}
	}

	m := Image{width, height, data[:]}
	pic.ShowImage(m)
}
