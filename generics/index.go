package main

import "fmt"
// [T comparable]がないと動作しない
// その型の値に対して == と != が使えるようになる
// [T any]だと == と != が使えない(下のようにエラーが出る)
// invalid operation: v == x (incomparable types in type set)
func Index[T any](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}

	return -1
}

func main() {
	si := []int{1, 2, 3, 4, 5}
	fmt.Println(Index(si, 3))

	ss := []string{"foo", "bar", "baz"}

	fmt.Println(Index(ss, "bar"))
}