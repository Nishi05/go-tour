package main

import "fmt"

type Vertex struct {
	Lat float64
}
// map は　キーと値のペアを保持するデータ構造
// map のゼロ値は nil で、nil マップはキーを持っていないため、キーを追加することはできない。
// make 関数は指定された型のマップを初期化し、使用可能な状態にする。
var m map[string]Vertex

func main() {
	// make 関数で初期化
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433,
	}
	fmt.Println(m["Bell Labs"])
}