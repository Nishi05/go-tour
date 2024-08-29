package main

import (
	"fmt"
	"math"
)

// Go言語にはClassの仕組みがない代わりに、型にメソッドを定義する事ができる
type Vertex struct {
	X, Y float64
}

// メソッドは、特別ねレシーバ引数を関数にとる。
// レシーバは、funcキーワードとメソッド名の間に自身の引数リストで表現される。
// この例では、Absメソッドは、Vertex型のレシーバvを持つ。
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}