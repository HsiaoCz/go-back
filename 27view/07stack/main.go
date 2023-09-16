package main

import (
	"errors"
	"log"
	"net/http"
	"sync"
	"time"
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

func (m *MStack) OutStack() (data any, err error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if len(m.data) == 0 {
		return nil, errors.New("当前栈为空")
	}
	data = m.data[len(m.data)-1]
	m.data = m.data[:len(m.data)-1]
	return data, err
}

const (
	addr = "127.0.0.1:9091"
)



func main() {
	http.HandleFunc("/user/set", HandleUserSet)
	srv := http.Server{
		Handler:      nil,
		Addr:         addr,
		ReadTimeout:  1500 * time.Millisecond,
		WriteTimeout: 1500 * time.Millisecond,
	}
	log.Fatal(srv.ListenAndServe())
}

func HandleUserSet(w http.ResponseWriter, r *http.Request) {

}
