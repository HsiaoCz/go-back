package queue

type IQueue interface {
	Push()
	Pop()
	Length()
	Size()
}

type Queue[T any] struct {
	node []T
}
