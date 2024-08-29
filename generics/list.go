package main

type List[T any] struct {
	next *List[T]
	val T
}

func (l *List[T]) Print() {
	for l != nil {
		println(l.val)
		l = l.next
	}
}

func main() {
	l := List[int]{val: 1}
	l.Print()
	l.next = &List[int]{val: 2}
	l.Print()
}