package main


import "fmt"

func main(){

	// defer キーワードを使うと、実行を呼び出し元の関数の終わり(returnする)まで遅延させることできる。
	// 引数自体はすぐに評価される。
	defer fmt.Println("world")

	fmt.Println("hello")
}