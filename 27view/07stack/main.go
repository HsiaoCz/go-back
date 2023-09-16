package main

import (
	"errors"
	"sync"
)

type MStack struct {
	lock   sync.Mutex
	data   []any
	length int
	maps   map[string]struct{}
}

func NewMStack(length int) *MStack {
	return &MStack{
		data:   make([]any, 0),
		length: length,
		maps:   make(map[string]struct{}),
	}
}

func (m *MStack) InStack(data any) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	if len(m.data) >= m.length {
		return errors.New("当前栈已满")
	}
	m.data = append(m.data, data)
	return nil
}

func main() {

}
