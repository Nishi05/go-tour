package main

import (
	"fmt"
	"math"
)


type Vertex struct {
	X, Y float64
}


func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


// 関数でも、ポインタを使って変数を変更することができる。
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}



func main() {
	v := Vertex{3, 4}
	// ポインタを渡す必要がある
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
