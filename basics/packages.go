package main // main関数を含む時にはmainパッケージを使う


// import文の書き方
import (
	"fmt"
	"math/rand" // 通常パッケージ名はインポートパスの最後の要素と同じ名前になる
)


func main(){
	fmt.Println("My favorite number is", rand.Intn(10))
}