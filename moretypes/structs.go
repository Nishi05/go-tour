package main

import "fmt"


// 構造体の宣言方法
type Vertex struct {
	X int
	Y int
}


func main(){
	fmt.Println(Vertex{1, 2})
}