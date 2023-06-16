package main

import "time"

// 用channel实现一个互斥锁
type MyMutex struct {
	ch chan struct{}
}

// 使用锁需要初始化
func NewMyMutex() *MyMutex {
	mu := &MyMutex{make(chan struct{}, 1)}
	mu.ch <- struct{}{}
	return mu
}

// 请求锁，直到获取到
func (m *MyMutex) Lock() {
	<-m.ch
}

// 解锁
func (m *MyMutex) UnLock() {
	select {
	case m.ch <- struct{}{}:
	default:
		panic("unlock of unlocked mutex")
	}
}

// 实现trylock，可轮询获取锁，如果成功，则返回true
// 如果失败，则返回 false
// 也就是说，这个方法无论成败都会立即返回，
// 获取不到锁（锁已被其他线程获取）时不会一直等待。

// 尝试获取锁
func (m *MyMutex) TryLock() bool {
	select {
	case <-m.ch:
		return true
	default:
	}
	return false
}

// trylock withtimeout
// tryLock(long, TimeUnit) - 可定时获取锁。
// 和 tryLock() 类似，区别仅在于这个方法在获取不到锁时会等待一定的时间，
// 在时间期限之内如果还获取不到锁，就返回 false。
// 如果如果一开始拿到锁或者在等待期间内拿到了锁，则返回 true。

func (m *MyMutex) TryLockWithTimeout(timeout time.Duration) bool {
	timer := time.NewTimer(timeout)
	select {
	case <-m.ch:
		timer.Stop()
		return true
	case <-timer.C:
	}
	return false
}
