package main


import "fmt"


func main(){
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// 型 []T は T のスライスを表し、固定長ではなく可変長。
	// [1:4]は最初の要素は含むが、最後の要素はのぞいた半開区間を選択する
	// なので primes[1] から primes[3] までが選択される
	var s []int = primes[1:4]

	fmt.Println(s)
}