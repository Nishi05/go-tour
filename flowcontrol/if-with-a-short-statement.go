package main

import (
	"fmt"
	"math"
)


func pow(x, n, lim float64) float64 {

	// for文のように、条件で評価する前に簡単なステートメントを書くことができる
	// 宣言された変数は ifのスコープ内でのみ有効
	if v := math.Pow(x, n); v < lim {
		return v
	}
	// vはここでは使えない
	// NG return v
	return lim
}

func main(){
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}