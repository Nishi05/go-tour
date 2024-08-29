package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	// 具体的な値としてnilを保持するinterfaceの値自体は、nilではない。
	describe(i) // (<nil>, *main.T)
	// interface自体の中にある具体的な値がnilの場合、メソッドはnilをレシーバとして呼び出す。
	i.M() // <nil>

	i = &T{"Hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}