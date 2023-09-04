package main

import "fmt"

// 代理模式
// 所谓的代理就是与被代理的对象具有相同的接口的类
// 客户端必须与目标类进行交互，而代理一般在交互的过程种，进行某些特别的处理

// 这里使用代购来实现
type Goods struct {
	Kind string // 商品的种类
	Fact bool   // 商品的真假
}

// 抽象的主题
type Shopping interface {
	Buy(goods *Goods) //代理的任务
}

// 实现层
// 具体的购物主题
type KoreaShopping struct{}

func (ks *KoreaShopping) Buy(goods *Goods) {
	fmt.Println("去韩国进行了购物,买了:", goods.Kind)
}

// 具体的购物主题
type AmericanShopping struct{}

func (as *AmericanShopping) Buy(goods *Goods) {
	fmt.Println("去美国进行了购物,买了:", goods.Kind)
}

// 设置代理
type OverseasProxy struct {
	shopping Shopping // 代理的某个主题
}

// 创建一个代理，配置关联被代理的主题
func NewProxy(shopping Shopping) Shopping {
	return &OverseasProxy{shopping: shopping}
}

// 代理实现的一些额外的功能
func (op *OverseasProxy) distinguish(goods *Goods) bool {
	fmt.Println("对[", goods.Kind, "]进行了辨别真伪.")
	if !goods.Fact {
		fmt.Println("发现假货", goods.Kind, ",不应该购买.")
	}
	return goods.Fact
}

func (op *OverseasProxy) check(goods *Goods) {
	fmt.Println("对[", goods.Kind, "]进行了海关检查,成功带回祖国")
}

func (op *OverseasProxy) Buy(goods *Goods) {
	// 1. 先验货
	if !op.distinguish(goods) {
		//2. 进行购买
		op.shopping.Buy(goods) //调用原被代理的具体主题任务
		//3 海关安检
		op.check(goods)
	}
}
