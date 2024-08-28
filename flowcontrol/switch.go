package main

import (
	"fmt"
	"runtime"
)

func main(){
	fmt.Print("Go runs on ")
	// for if 同様、switchでも変数を宣言できる(変数はswitchのスコープ内でのみ有効)
	// 他の言語と違い、caseの最後にbreakを書く必要はない
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

}