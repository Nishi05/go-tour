package main

import "fmt"
// factored 宣言
// 型のない定数は、任意の精度の値を取る
// Bigはint型だとオーバーフローするが、float64型だとオーバーフローしない
// Smallは intでもfloat64でも使える
const (
	Big = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main(){
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	fmt.Println(needInt(Big))
}
