package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// トップレベルに指定した型が単純である場合には、型名をリテラルから省略できる
var m = map[string]Vertex {
	"Bell Labs": {40.68433, -74.39967},
	"Google": {37.42202, -122.08408},
}


func main(){
	fmt.Println(m)
}