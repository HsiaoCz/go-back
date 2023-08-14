package main

// go context
// context，中文译作“上下文”，
// 准确说它是 goroutine 的上下文，
// 包含 goroutine 的运行状态、环境、现场等信息。
// context 主要用来在 goroutine 之间传递上下文信息，
// 包括：取消信号、超时时间、截止时间、k-v 等

// context的使用规范
// 不要将 Context 放入结构体，相反 Context 应该作为第一个参数传入，命名为 ctx
// 即使函数允许，也不要传入 nil 的 Context。如果不知道用哪种 Context，可以使用context.TODO()
// 使用 Context 的 Value 相关方法只应该用于在程序和接口中传递的和请求相关的元数据，不要用它来传递一些可选的参数
// 相同的 Context 可以传递给在不同的 goroutine，因为 Context 是并发安全的

/**
   context 本身是一个接口
   type Context interface{
	    // 返回一个channel。当times out或者调用cancel方法时，将会close掉。
    	Done()<-chan struct{}
		// 返回一个错误。该context为什么被取消掉
		Err()error
		//  返回截止时间和 ok
		Deadline()(deadline time.Time,ok bool)
		// 返回 Key 值
		Value(key interface{})interface{}
   }

   context 包的方法
     // 这两个方法都会返回一个空的context  
     func Background() Context
     func TODO() Context
     
	 //  以一个新的 Done channel 返回一个父 Context 的拷贝
     func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
	 //  的最后期限调整为不晚于 deadline 返回父上下文的副本
     func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
	 //  返回 WithDeadline(parent, time.Now().Add(timeout))
     func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
     //  返回的父与键关联的值在 val 的副本   
	 func WithValue(parent Context, key, val interface{}) Context

**/

func main() {

}
