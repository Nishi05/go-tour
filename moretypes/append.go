package main

import "fmt"

// append 関数を使用すると、スライスへ新しい要素を追加できる
// スライスの容量が足りない場合は、元の配列の容量を2倍に拡張して新しい配列を割り当てる
// (2倍でも足りない場合は、元の配列の容量に対してx3, x4 と拡張される)

func main(){
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)


	s = append(s, 2, 3, 4, 5, 6)
	printSlice(s)

	s = append(s, 3, 4)
	printSlice(s)
}

func printSlice(s []int){
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}