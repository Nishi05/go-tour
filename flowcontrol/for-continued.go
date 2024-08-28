package main

import "fmt"

func main(){
	sum := 1
	// 初期化と後処理ステートメントは省略可能
	// for ; sum < 1000; sum += sum {} と書くこともできる(回答は一緒)
	for ; sum < 1000; {
		sum += sum
	}

	fmt.Println(sum)
}