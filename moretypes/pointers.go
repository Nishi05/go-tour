package main


import "fmt"


func main(){
	i, j := 42, 2701


	p := &i // i のアドレスを代入
	fmt.Println(*p) // ポインタ p を通して i の値を読む
	*p = 21 // ポインタ p を通して i に値を代入
	fmt.Println(i) // i の値を表示

	p = &j // j のアドレスを代入
	*p = *p / 37 // ポインタ p を通して j の値を割り算
	fmt.Println(j) // j の値を表示
}