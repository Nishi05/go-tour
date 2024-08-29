package main


import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

// for文で利用できる range は、スライスやマップを一つずつ反復処理するために使います。
// その際に、range は反復毎に2つの値を返します。1つ目の値がインデックスで、2つ目がそのインデックスの要素のコピー。
func main() {
	// range 〇〇 の 〇〇 にスライスやマップを指定する
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}