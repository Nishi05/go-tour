package main

import "fmt"


type Vertex struct {
	X, Y int
}

//　構造体では、フィールドの値を指定しない場合、ゼロ値が入る
var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1} // Y:0 が暗黙的に入る
	v3 = Vertex{} // X:0 と Y:0 が暗黙的に入る
	p = &Vertex{1, 2} // 構造体へのポインタ
	p2 = &v3 // 構造体変数へのポインタ
)


func main(){
	p2.X = 3 // v3.X に 3 を代入
	fmt.Println(v1, p, v2, v3)
}