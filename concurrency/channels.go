package main

import "fmt"

func sum(s []int, c chan int){
	sum := 0
	for _, v := range s{
		sum += v
	}
	c <-sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	// channel型は、チャンネルオペレータ <- を用いて値の送受信ができる通り道。（データは矢印の方向に流れる）
	// make 関数を用いてチャンネルを作成する
	c := make(chan int)
	// 作業を分担して、並行処理を行う
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}