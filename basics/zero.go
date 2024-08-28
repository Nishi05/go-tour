package main

import "fmt"

// 初期化しない場合は、ゼロ値(全てが0ではない)が入る
func main() {
	var i int // 0
	var f float64 // 0
	var b bool // false
	var s string // "" (空文字列)
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}