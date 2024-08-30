package main


import (
	"fmt"
	"time"
)


func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	// ブロックせずに送受信をするなら、defaultのcaseを使う
	// default は、他のcaseが準備できない場合に実行される
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}