package main

import "fmt"

func main() {
	var i interface{} = "hello"

	// 型アサーション
	// interfaceの値の基になる具体的な名前を利用する手段
	// mapから値を取り出すときと似ている
	s := i.(string)
	fmt.Println(s)

	// インターフェースの値が特定の型を保持しているかどうかをテストするために、型アサーションは2つの値を返す。
	// 1つ目はアサーションした値、2つ目はアサーションが成功したかどうかを示すブール値。
	s, ok := i.(string)
	fmt.Println(s, ok)
	
	
	// iがfloat64を保持していない場合、fにはfloat64のゼロ値が入る。(つまり f = 0, ok = false)
	f, ok := i.(float64)
	fmt.Println(f, ok)

	// テスト結果を代入する変数がない場合に、型アサーションが失敗するとpanicが発生する。
	// ランタイムエラーになる
	f = i.(float64) // panic
	fmt.Println(f)
}