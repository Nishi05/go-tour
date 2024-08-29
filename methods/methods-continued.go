package main


import (
	"fmt"
	"math"
)

// struct(構造体)だけではなく、任意の型にもメソッドを定義できる
type MyFloat float64


func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}


func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

/*
	ただし、レシーバの型が同じパッケージ内にある必要がある。
	他のパッケージの型に対してメソッドを定義することはできない。
	
	組み込みの型(intなど)にもメソッドは定義できない。
*/