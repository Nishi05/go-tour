package main

import "fmt"

// スライスするときは、それらの既定値を代わりに使用することで上限または下限を省略できる
// 既定値は下限が 0 で上限がスライスの長さ
func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s) // 3, 5, 7

	s = s[:2]
	fmt.Println(s) // 3, 5

	s = s[:]
	fmt.Println(s) // 3, 5
}

/*
var a [10]int は下記と同じ

a[0:10]
a[:10]
a[0:]
a[:]
*/