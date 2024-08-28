package main

import "fmt"

// イテレーション = 繰り返し
func main(){
	sum := 0
	// for キーワードが必要で () は不要
	// i := 0 最初のイテレーションが始まる前に実行される
	// i < 20 イテレーションごとに評価される
	// i++ イテレーションごとの最後に実行されう
	for i := 0; i < 20; i++ {
		sum += i
	}
	fmt.Println(sum)
}