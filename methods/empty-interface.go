package main

import "fmt"

func main() {
	// ゼロ個のメソッドを指定されたインターフェース型は、空のインターフェースと呼ばれる。
	//　空のインターフェースは任意の型の値を保持できる。(全ての型は、少なくともゼロ個のメソッドを実装しているため)
	var i interface{}
	describe(i)

	i = 42
	describe(i) // (42, int)

	i = "hello"
	describe(i) // (hello, string)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}