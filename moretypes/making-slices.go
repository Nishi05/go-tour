package main

import "fmt"

// make 関数を使用してスライスを作成する(動的サイズの配列を作成する方法)
// make 関数はゼロ化された配列を割り当て、その配列を指すスライスを返す
// ex) a := make([]int, 5) lent = 5 cap = 5 [0 0 0 0 0]
// make 関数の3番目の引数には、スライスの容量を指定できる
func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)
	// 予想　c len=2 cap5 [0 0]

	d := c[2:5]
	printSlice("d", d)
	// 予想 d len=3 cap=3 [0 0 0]

}



func printSlice(s string, x []int ) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}