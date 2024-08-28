package main

// パッケージの factored 宣言
import (
	"fmt"
	"math/cmplx"
)

// 変数の factored 宣言
var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type %T Value %v\n", ToBe, ToBe)
	fmt.Printf("Type %T Value %v\n", MaxInt, MaxInt)
	fmt.Printf("Type %T Value %v\n", z, z)
}

/* 
色々な型が存在する

bool 

string

int ini8 int16 int32 int64 // 符号付き整数(マイナスも対応可能)
uint uint8 uint16 uint32 uint64 uintptr　// 符号なし整数 (マイナスは対応しない)

byte // uint8 の別名

rune // int32 の別名
		 // Unicode のコードポイントを表す

float32 float64 // 浮動小数点にも対応

complex64 complex128 // 複素数にも対応

// 参考
// https://qiita.com/kyosuke5_20/items/8a3af2411f79a25f78b5

*/