package main

import "fmt"

// 関数の外では、暗黙的な型宣言はできない
// NG o := 42

func main() {
	var i, j int = 1, 2
	k := 3 // var 宣言の代わりに := の代入分を使い、暗黙的な型宣言ができる
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}