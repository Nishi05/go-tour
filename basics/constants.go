package main

import "fmt"
// 定数を宣言するには const キーワードを使う
// 文字列、boolean、数値が定数で宣言できる
// := を使って宣言はできない
// 型は指定できる const Pi float32 = 3.14
const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

}