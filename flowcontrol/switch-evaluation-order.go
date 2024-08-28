package main

import (
	"fmt"
	"time"
)


func main(){
	fmt.Println("When's Saturday?")

	today := time.Now().Weekday()
	// 上から下に評価され、一致すればそのcaseの処理が実行され停止する(自動的にbreakされる)
	// caseの値は定数である必要がなく変数でも良い(整数である必要もない)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	case today + 3:
		fmt.Println("In three days.")
	default:
		fmt.Println("Too far away.")
	}
}