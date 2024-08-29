package main

import "fmt"


// 引数としても返り値としても関数を利用する事ができる
func compute(fn func(int, int) int) int {
	return fn(3, 4)
}

func main() {
	// 変数へ関数を渡すこともできる
	hypot := func(x, y int) int {
		return x + y
	}

	fmt.Println(hypot(5, 13))

	fmt.Println(compute(hypot))
}