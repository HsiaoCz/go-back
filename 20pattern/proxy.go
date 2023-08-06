package main

import "fmt"

// 代理模式
// 代理和被代理的对象有相同的接口，客户端只能通过代理与目标类进行交互
// 而代理一般在交互的过程中，进行一些特别的处理

type Goods struct {
	Kind string
	Fact bool
}

// 抽象的主题
type Shopping interface {
	Buy(goods *Goods) // 代理的任务
}

// 实现层，具体的购物主题
type KoreaShopping struct{}

func (k *KoreaShopping) Buy(goods *Goods) { fmt.Println("去韩国购物,买了:", goods.Kind) }

type AmericanShopping struct{}

func (as *AmericanShopping) Buy(goods *Goods) { fmt.Println("去美国进行了购物...", goods.Kind) }

// 设置代理
type OverseasProxy struct {
	shopping Shopping
}

func NewProxy(shopping Shopping) Shopping {
	return &OverseasProxy{shopping: shopping}
}

func (o *OverseasProxy) Buy(goods *Goods) {
	if o.distinguish(goods) {
		o.shopping.Buy(goods)
		o.check(goods)
	}
}

func (o *OverseasProxy) distinguish(goods *Goods) bool {
	fmt.Println("对货物进行了验货...")
	if !goods.Fact {
		fmt.Println("发现假货:", goods.Kind)
	}
	return goods.Fact
}
func (o *OverseasProxy) check(goods *Goods) {
	fmt.Println("对货物进行了验货")
}
