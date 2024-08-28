package main

import "fmt"


// 配列の長さも型の一部なので、長さを変えることはできない
func main(){
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)


	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}