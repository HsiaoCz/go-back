package main

import "fmt"

// 代理模式
// 这里使用代购来实现

type Goods struct {
	Kind string // 商品的种类
	Fact bool   // 商品的真假
}

// 抽象的主题
type Shopping interface {
	Buy(goods *Goods)
}

// 实现层
// 具体的购物主题

type KoreaShopping struct{}

func (ks *KoreaShopping) Buy(goods *Goods) {
	fmt.Println("去韩国进行了购物,买了:", goods.Kind)
}

type AmericanShopping struct{}

func (as *AmericanShopping) Buy(goods *Goods) {
	fmt.Println("去美国进行了购物,买了:", goods.Kind)
}

// 设置代理
type OverseasProxy struct {
	shopping Shopping
}

func NewProxy(shopping Shopping) Shopping {
	return &OverseasProxy{shopping: shopping}
}

func (o *OverseasProxy) distinguish(goods *Goods) bool {
	fmt.Printf("对%s进行了辨别真伪", goods.Kind)
	if !goods.Fact {
		fmt.Printf("发现假货:%s不应该购买", goods.Kind)
	}
	return goods.Fact
}

func (o *OverseasProxy) check(goods *Goods) {
	fmt.Printf("对%s进行了海关检查,成功带回祖国", goods.Kind)
}

func (o *OverseasProxy) Buy(goods *Goods) {
	if o.distinguish(goods) {
		o.shopping.Buy(goods)
		o.check(goods)
	}
}

func main() {
	g1 := Goods{
		Kind: "韩国面膜",
		Fact: true,
	}

	g2 := Goods{
		Kind: "日本马桶盖",
		Fact: false,
	}
	ovProxy := NewProxy(new(KoreaShopping))
	ovProxy.Buy(&g1)
	ovProxy.Buy(&g2)
}
