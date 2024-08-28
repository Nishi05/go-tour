package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!" //初期値を指定すると、型を省略する事ができ、変数は初期値の型になる

	fmt.Println(i, j, c, python, java)
}