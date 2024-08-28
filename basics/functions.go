package main


import "fmt"

// functionsは0個以上の引数を取ることができる
// 変数名の後に型名を書く
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}