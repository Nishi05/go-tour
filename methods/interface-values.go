package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Printf("(%v, %T)\n", t, t)
	fmt.Println(t.S)
}


type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	// interfaceの値は、値と具体的な型のタプルのように考える事ができる。
	var i I

	i = T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

/*
インターフェースの値がポインタの場合
レシーバはポインタと値どちらでも呼び出す事ができる。
インターフェースの値が値の場合
レシーバは値のみ呼び出す事ができる。
*/