package main


import "fmt"


func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
// Goのチャネルにおいて、バッフォサイズが0の場合受信操作が行われるまで送信操作も完了しないという特性がある。
func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			// 受信操作(<-c)が行われるまで、送信操作(c <- x)は完了しない
			// なので、フィボナッチ数列の値が最初の10個というのが保証される
			fmt.Println(<-c)
		}
		// なので、11個目を受信することもできる。
		fmt.Println(<-c) // 55
		quit <- 0
	}()
	fibonacci(c, quit)
}