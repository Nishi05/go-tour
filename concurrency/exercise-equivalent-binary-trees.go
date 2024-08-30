package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)


func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := range ch1 {
		if i != <-ch2 {
			return false
		}
	}

	return true
}



func walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	walk(t.Left, ch)
	ch <- t.Value
	walk(t.Right, ch)
}


func main () {
	// ch := make(chan int)
	// go Walk(tree.New(1), ch)
	// for i := range ch {
	// 	println(i)
	// }

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(2), tree.New(2)))
	fmt.Println(Same(tree.New(3), tree.New(3)))

	fmt.Println(Same(tree.New(1), tree.New(2)))
	fmt.Println(Same(tree.New(1), tree.New(3)))
	fmt.Println(Same(tree.New(1), tree.New(4)))

	fmt.Println(Same(tree.New(2), tree.New(1)))
	fmt.Println(Same(tree.New(3), tree.New(1)))
	fmt.Println(Same(tree.New(4), tree.New(1)))
}