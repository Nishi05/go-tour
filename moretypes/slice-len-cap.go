package main

import "fmt"


// スライスの長さ＝そうに含まれる要素数 len()で取得
// スライスの容量＝スライスの最初の要素から数えて、元となる配列の要素数 cap()で取得
// 容量があるかぎりスライスの長さを増やすことができる

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	// s = s[2:]だと、直前にスライスした長さまでしか取得しない(つまり s[2:4])
	s = s[2:6] // 5, 7, 11, 13
	printSlice(s)

}


func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}