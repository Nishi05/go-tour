package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main(){

	v := Vertex{1, 2}
	// 構造体のフィールドへのアクセスは、構造体変数名.フィールド名で行う
	v.X = 4

	fmt.Println(v.X)
}