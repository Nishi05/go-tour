package main


import (
	"fmt"
	"math"
)

// interface型はメソッドのシグネチャの集まりで定義される。
type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	// interface型に定義されているメソッドを実装した値を、interface型の変数に代入することができる。
	a = f
	a = &v

	// Vertex型のAbs()メソッドは*Vertex型に定義されているため、Vertex型の変数をinterface型の変数に代入することはできない。
	// a = v

	fmt.Println(a.Abs())
}