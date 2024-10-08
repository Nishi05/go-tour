package main

import (
	"fmt"
	"math"
)


type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}


/*
ポインタレシーバを使う理由
1. メソッドがレシーバを指す先の変数を変更するため
2. メソッドの呼び出し毎に変数のコピーを避けるため(レシーバが大きな構造体の時に効果的)

一般的に、変数レシーバ、または、ポインタレシーバのどちらかで全てのメソッドを使うべき。(混在させるべきはない)
*/

