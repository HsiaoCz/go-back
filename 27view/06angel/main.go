package main

import "sync"

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
	m.data = append(m.data, data)
	return nil
}

func main() {

}
