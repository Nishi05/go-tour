package main

import "fmt"

/*
by ChatGPT
クロージャー（Closure） とは、関数とその関数が定義された環境（スコープ）を合わせ持つものを指す。
クロージャーは、関数が外部の変数を参照し、その変数の状態を保持または変更することができる機能を提供する
*/

// この例では、adder 関数は、クロージャーを返す。
// 関数が自身の外部にある変数をキャプチャ（捕捉）し、その変数へのアクセスを持ち続けることができる。
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}


func main(){
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}