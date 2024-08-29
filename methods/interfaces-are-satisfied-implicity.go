package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}
// interface を実装することを明示的に宣言する必要はない。(implementsキーワードは不要)
func (t T) M() {
	fmt.Println(t.S)
}


func main() {
	var i I = T{"hello"}
	i.M()
}