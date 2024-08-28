package main

import "fmt"

// スライスのゼロ値は nil です。
// nil スライスは長さと容量は０で, 何の元となる配列も持っていません。
func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}