package main

import "fmt"

// 関数の返り値に名前をつける事ができる。
// 名前をつけた戻り値を使うと、明示的にreturn文に返す変数を指定しなくても良い(naked return と呼ばれる)
// ただし、短い関数でのみnaked returnは使うべきである。長い関数では読みやすさに悪影響を及ぼす。
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}