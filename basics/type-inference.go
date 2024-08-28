package main

import "fmt"

// 型を指定していない場合は、右辺の型から型推論される
// 右辺が型をもっている場合は、左辺の型は同じになる
func main() {
	v := 42.0
	w := v

	fmt.Printf("v is of type %T %T\n", v, w)
}