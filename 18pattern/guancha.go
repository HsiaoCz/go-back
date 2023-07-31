package main

import "fmt"

// 观察者模式
// 观察者模式用于建立一种对象之间的依赖关系，一个对象发生改变时将通知其他对象
// 其他对象将作出相应的反应

// 抽象的观察者
type Listener interface {
	OnTeacherComming() // 观察者得到通知后要做的动作
}

type Notifier interface {
	AddListener(listener Listener)
	RemoveListener(listener Listener)
	Notify()
}

// 实现层
// 观察者学生
type StuZhang struct {
	BadThing string
}

func (s *StuZhang) OnTeacherComming() { fmt.Println("张三 停止了:", s.BadThing) }
func (s *StuZhang) DoBadThing()       { fmt.Println("张三正在:", s.BadThing) }

type StuZhao4 struct {
	BadThing string
}

func (s *StuZhao4) OnTeacherComming() { fmt.Println("赵四 停止了:", s.BadThing) }
func (s *StuZhao4) DoBadThing()       { fmt.Println("赵四正在:", s.BadThing) }

type StuWang5 struct {
	BadThing string
}

func (s *StuWang5) OnTeacherComming() { fmt.Println("王五 停止了:", s.BadThing) }
func (s *StuWang5) DoBadThing()       { fmt.Println("王五正在:", s.BadThing) }

// 通知者班长
type ClassMonitor struct {
	listenerLsit []Listener
}

func (m *ClassMonitor) AddListener(listener ...Listener) {
	m.listenerLsit = append(m.listenerLsit, listener...)
}

func (m *ClassMonitor) RemoveListener(listeners ...Listener) {
	for _, list := range listeners {
		for index, mlist := range m.listenerLsit {
			if list == mlist {
				m.listenerLsit = append(m.listenerLsit[:index], m.listenerLsit[index+1:]...)
				break
			}
		}
	}
}

func (m *ClassMonitor) Notifier() {
	for _, listener := range m.listenerLsit {
		listener.OnTeacherComming()
	}
}

func main() {
	s1 := &StuZhang{
		BadThing: "抄作业",
	}
	s2 := &StuZhao4{
		BadThing: "玩王者荣耀",
	}
	s3 := &StuWang5{
		BadThing: "看赵四玩王者荣耀",
	}

	classMonitor := new(ClassMonitor)

	fmt.Println("上课了，但是老师没有来，学生们都在忙自己的事...")
	s1.DoBadThing()
	s2.DoBadThing()
	s3.DoBadThing()

	classMonitor.AddListener(s1, s2, s3)
	fmt.Println("这时候老师来了,班长给学生使了一个眼神.......")
	classMonitor.Notifier()

	classMonitor.RemoveListener(s1, s2)
	fmt.Println("....................")
	classMonitor.Notifier()
}
