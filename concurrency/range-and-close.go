package main


import "fmt"

// 送り手
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// close すると、チャンネルにこれ以上送信することができなくなる
	// 受け手のrangeを終了させたい場合に使う（頻繁には使用しない）
	close(c)
}

func main() {
	c := make(chan int, 10)

	go fibonacci(cap(c), c)
	// channelがcloseされるまで、値を受信し続ける
	// 受け手
	for i := range c {
		fmt.Println(i)
		// 途中で受け手が終了させるとpanic(エラー)が発生する
		//　NG if i == 13 {
		// 	close(c)
		// }
	}
}