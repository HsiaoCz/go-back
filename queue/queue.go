package queue

import "sync"

type IQueue interface {
	Push()
	Pop()
	Length()
	Size()
}

type Queue[T any] struct {
	lock sync.Mutex
	node []T
}

func (a Queue[T]) Push(data T) {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.node = append(a.node, data)
}
