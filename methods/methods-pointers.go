package main

import (
	"fmt"
	"math"
)


type Vertex struct {
	X, Y float64
}


func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
// ポインタレシーバでメソッドを宣言できる。
// レシーバが指す変数を変更する際に使用する事が多い。(より一般的に使われる)
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}


func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}