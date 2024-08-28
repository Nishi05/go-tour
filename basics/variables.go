package main

import "fmt"

// 変数の宣言は関数の引数や返り値と同じように複数の変数の最後に書くことでまとめて宣言できる
// var ステートメントはパッケージ内や関数で利用でき、変数を宣言するために使用できる
var c, python, java bool

func main(){
	var i int
	fmt.Println(i, c, python, java)
}
