package main

import (
	"image"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {

	slice_y := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		slice_x := make([]uint8, dx)

		for x := 0; x < dx; x++ {
			slice_x[x] = uint8((x + y) / 2)
		}

		slice_y[y] = slice_x
	}

	return slice_y
}

func ShowImage(m image.Image) {
	ShowImage(m)
}

func main() {
	pic.Show(Pic)
}
