package main

import "fmt"

func main(){
	pow := make([]int, 10)
	// インデックスだけを使う場合は、値を省略できる。
	for i := range pow {
		pow[i] = i << uint(i)
	}
	// _ (アンダーバー) は、インデックスや値を使わない場合に指定できる
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}