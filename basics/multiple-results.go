package main

import "fmt"


// 関数の戻り値を複数指定する事ができる
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}