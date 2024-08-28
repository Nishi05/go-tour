package main

import (
	"fmt"
	"math"
)
// C言語とは異なり、Go言語では明示的な型変換が必要
func main() {
	var x, y int = 3, 4 // x, y := 3, 4 と同じ
	var f float64 = math.Sqrt(float64(x*x + y*y)) // int -> float64 に型変換をしないとエラーになる
	var z uint = uint(f) // 型変換　z := uint(f)でもOK
	fmt.Println(x, y, z)
}