package pattern

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 单例的饿汉式
type sinelehan struct{}

var instanceeh *sinelehan

var initialize uint32

var lock sync.Mutex

func GetInstanceEH() *sinelehan {
	if atomic.LoadUint32(&initialize) == 1 {
		return instanceeh
	}
	lock.Lock()
	defer lock.Unlock()
	if initialize == 0 {
		instanceeh = new(sinelehan)
		atomic.StoreUint32(&initialize, 1)
	}
	return instanceeh
}

func (s *sinelehan) DoSomeFun() { fmt.Println("单例的某个方法") }
