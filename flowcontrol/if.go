package main

import (
	"fmt"
	"math"
)
// if文もfor文同様に () は不要
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	} 

	return fmt.Sprint(math.Sqrt(x))
}

func main(){
	fmt.Println(sqrt(2), sqrt(-4))
}