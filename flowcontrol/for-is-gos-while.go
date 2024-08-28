package main

import "fmt"

func main(){
	sum := 1
	// ; を省略すると C言語の while 文と同じような使い方ができる
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}