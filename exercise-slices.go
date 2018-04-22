package main

import (
	//"fmt"
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) (s [][]uint8) {
	for y := 0; y < dy; y++ {
		var ss []uint8
		for x := 0; x < dx; x++ {
			var value int
			//value := (x + y) / 2
			/*
			   if (x + y) % 2 == 0 {
			           value = 0
			                 } else {
			                         value = 255
			                               }
			*/
			//value = (x * y)
			value = x ^ y
			ss = append(ss, uint8(value))
		}
		s = append(s, ss)
	}
	return s
}

func main() {
	pic.Show(Pic)
}
