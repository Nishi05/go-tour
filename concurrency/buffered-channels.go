package main

import "fmt"

func main() {
	ch := make(chan int, 2) // 二つ目の引数でバッファサイズを指定できる
	ch <- 1
	ch <- 2
	ch <- 3 // バッファ以上のデータを送信すると、エラーが出る
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}