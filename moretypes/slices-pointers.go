package main

import "fmt"


// スライスは、配列への参照のようなもの
func main() {

	names := [4]string{
		"A",
		"B",
		"C",
		"D",
	}
	fmt.Println(names)
	// names[0:2]を代入しようとすると、基本は可変長になる
	// 固定長にしたい場合は、以下のようにする
	// var a [2]string =  [2]string(names[0:2])
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	a[1] = "YYY"
	fmt.Println(a, b)

	fmt.Println(names)
}