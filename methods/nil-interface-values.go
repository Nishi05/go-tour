package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i) // (<nil>, <nil>)
	// 呼び出す具体的なメソッドを示す型がインターフェースのタプルの中に存在しない場合、メソッドを呼び出すとランタイムエラーになる。
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}