package main

import "fmt"

func do(i interface{}) {
	// 型 switchと呼ばれる構文
	// caseでは型を指定する
	// この例では、iの型によって異なる処理を行っている
	// 宣言時には、型アサーションと同じ構文を持つが、キーワードはtypeを指定する。
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}