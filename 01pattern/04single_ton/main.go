package main

import "fmt"

// 单例模式
// 单例保证一个类永远只有一个对象，且该对象的功能依然能被其他模块所使用
// 单例有三个要点
// 1.某个类只能有一个实例
// 2.它必须自行创建这个实例
// 3.他必须自行向整个系统提供这个实例

// 保证一个类永远只有一个对象
// 1.保证这个类非共有化，外界不能通过这个类直接创建一个对象
// 那么这个类应该是非公有访问的，类首字母要小写

type singleton struct{}

// 2.这个类需要向外提供一个实例
// 那么需要一个指针指向这个类，而且指针不会改变方向

var instance *singleton = new(singleton)

// 3.如果全部都是私有化的，那么外部模块永远无法访问这个类和对象
// 所以需要对外提供一个唯一的方法来获取这个唯一实例对象
// 注意：这个方法是否可以定义为singelton的一个成员方法呢？
// 答案是不能，成员方法需要先访问类，才能访问到方法
func GetInstance() *singleton {
	return instance
}

func (s *singleton) SomeThing() {
	fmt.Println("单例对象的某些方法")
}

func main() {

	// 这种单例是饿汉式单例
	// 在初始化单例唯一的指针的时候，就已经提前开辟好了一个对象，申请了内存
	// 饿汉式的好处是，不会出现线程并发创建，出现多个单例
	// 缺点是即使业务没有被使用，也始终申请了一块内存
	s := GetInstance()
	s.SomeThing()
}
