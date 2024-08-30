package main

import (
	"fmt"
	"time"
)


func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(time.Millisecond)
		fmt.Println(s)
	}
}

// goroutineは、Goのランタイムに管理される軽量なスレッド。
// 非同期的な処理を実現する。(多分)
func main() {
	go say("world")
	say("hello")
}