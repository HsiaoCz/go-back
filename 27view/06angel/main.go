package main

import (
	"errors"
	"sync"
)

type Book struct {
	Identity string `json:"identity"`
	BookName string `json:"book_name"`
	Title    string `json:"title"`
	Auther   string `json:"auther"`
	Type     string `json:"type"`
	Content  string `json:"content"`
}

type MQueue struct {
	data  []Book
	lock  sync.Mutex
	lenth int
	maps  map[string]struct{}
}

func NewMQueue(lenth int) *MQueue {
	return &MQueue{
		data:  make([]Book, 0),
		lenth: lenth,
		maps:  make(map[string]struct{}),
	}
}

func (m *MQueue) InMQueue(data Book) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	_, ok := m.maps[data.Identity]
	if ok {
		return errors.New("当前数据在队列中已经存在,请勿重复添加")
	}
	if len(m.data) == m.lenth {
		return errors.New("当前队列已满,请稍后添加")
	}
	m.data = append(m.data, data)
	m.maps[data.Identity] = struct{}{}
	return nil
}

func (m *MQueue) OutMQueue() (book Book, err error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if len(m.data) == 0 {
		return Book{}, errors.New("当前队列为空")
	}
	book = m.data[0]
	m.data = m.data[1:]
	delete(m.maps, book.Identity)
	return book, nil
}

func main() {

}
