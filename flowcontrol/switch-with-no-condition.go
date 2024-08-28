package main

import (
	"fmt"
	"time"
)


func main(){
	t := time.Now()
	// switch文に条件を書かない場合は、"if then else"のように使えるかつシンプルに表現できる
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	/* 
	// 以下のようにif then elseを使っても同じ結果が得られるが、switch文の方がシンプルである
	if t.Hour() < 12 {
		fmt.Println("Good morning!")
	} else if t.Hour() < 17 {
		fmt.Println("Good afternoon.") 
	} else {
	 	fmt.Println("Good evening.") 
	}
	*/
}