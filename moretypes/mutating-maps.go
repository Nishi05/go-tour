package main

import "fmt"

func main(){
	m := make(map[string]int)


	// 挿入
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	// 更新
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// 削除
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])


	// vは要素の取得 ない場合は要素の型のゼロ値が返される(この場合はintなので0)
	// okはそのキーが存在するかどうかを示す
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
} 