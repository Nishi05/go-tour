package main 

import "fmt"


func main(){
	fmt.Println("counting")


	// defer へ渡した関数はstackされる。
	// 呼び出しの順番はLIFO(last-in-first-out)である。
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}