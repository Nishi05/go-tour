package main

import "golang.org/x/tour/pic"


func Pic(dx, dy int) [][]uint8 {
	pic :=make([][]uint8, dy)

	for y := range pic {
		pic[y] = make([]uint8, dx)
		for i := range pic[y] {
			pic[y][i] = uint8((i + y) / 2)
		}
	}
	return pic
}


func main() {
	pic.Show(Pic)
}