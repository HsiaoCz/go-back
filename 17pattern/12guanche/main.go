package main

import "fmt"

// 观察者模式
// 观察者模式用于建立一种对象之间的依赖关系，一个对象发生改变时将通知其他对象
// 其他对象将作出相应的反应
// 这里举个例子

// 抽象的观察者
type Listener interface {
	OnTeacherComming() // 观察者得到通知后要触发的动作
}

type Notifier interface {
	AddListener(listener Listener)
	RemoveListener(listener Listener)
	Notify()
}

// 实现层
// 观察者学生
type StuZhang struct{ BadThing string }

func (s *StuZhang) OnTeacherComming() { fmt.Println("张三 停止了..", s.BadThing) }
func (s *StuZhang) DoBadThing()       { fmt.Println("张三 正在", s.BadThing) }

type StuZhao4 struct{ Badthing string }

func (s *StuZhao4) OnTeacherComming() { fmt.Println("赵4 停止 ", s.Badthing) }

func (s *StuZhao4) DoBadthing() { fmt.Println("赵4 正在", s.Badthing) }

type StuWang5 struct{ Badthing string }

func (s *StuWang5) OnTeacherComming() { fmt.Println("王5 停止 ", s.Badthing) }

func (s *StuWang5) DoBadthing() { fmt.Println("王5 正在", s.Badthing) }

// 通知者班长
type ClassMonitor struct{ listenerList []Listener }

func (m *ClassMonitor) AddListener(listener Listener) {
	m.listenerList = append(m.listenerList, listener)
}

func (m *ClassMonitor) RemoveListener(listener Listener) {
	for index, l := range m.listenerList {
		// 找到要删除的元素位置
		if listener == l {
			// 将删除的点前后的元素链接起来
			m.listenerList = append(m.listenerList[:index], m.listenerList[index+1:]...)
			break
		}
	}
}

func (m *ClassMonitor) Notify() {
	for _, listener := range m.listenerList {
		//依次调用全部观察的具体动作
		listener.OnTeacherComming()
	}
}

func main() {
	s1 := &StuZhang{
		BadThing: "抄作业",
	}
	s2 := &StuZhao4{
		Badthing: "玩王者荣耀",
	}
	s3 := &StuWang5{
		Badthing: "看赵四玩王者荣耀",
	}

	classMonitor := new(ClassMonitor)

	fmt.Println("上课了，但是老师没有来，学生们都在忙自己的事...")
	s1.DoBadThing()
	s2.DoBadthing()
	s3.DoBadthing()

	classMonitor.AddListener(s1)
	classMonitor.AddListener(s2)
	classMonitor.AddListener(s3)

	fmt.Println("这时候老师来了，班长给学什么使了一个眼神...")
	classMonitor.Notify()
}
