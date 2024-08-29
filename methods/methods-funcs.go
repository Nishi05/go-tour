package main

import (
	"fmt"
	"math"
)


type Vertex struct {
	X, Y float64
}

// メソッドではなく、通常通りに関数を定義する場合
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}



func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}