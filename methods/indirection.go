package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}


func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}


func main(){
	v := Vertex{3, 4}

	// ポインタレシーバは利便性のため明示的にポインタを指定する必要がない。
	// ((&v).Scale(2)と自動的に解釈される)
	v.Scale(2)
	ScaleFunc(&v, 10)


	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)


	fmt.Println(v, p)
}