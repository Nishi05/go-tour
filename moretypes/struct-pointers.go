package main

import "fmt"

type Vertex struct {
	X int
	Y int
}


func main(){
	v := Vertex{1, 2}
	p := &v
	// 構造体のフィールドへのアクセスは、構造体のポイントを通してもアクセスできる
	(*p).X = 1e9 // p.X と書いても良い（(*p)と書くのが面倒だから）

	fmt.Println(v.X)
}