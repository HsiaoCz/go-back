package main

import (
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

// 如何处理并发错误
// errGroup
// type Group struct {
// 	cancel func()

// 	wg sync.WaitGroup

// 	errOnce sync.Once
// 	err     error
// }

// errgroup提供了Go和Wait两个方法
// func (g *Group) Go(f func() error)
// go函数会在新goroutine中调用传入的函数f
// 第一个返回非零错误的调用将取消该Group；下面的Wait方法会返回该错误

// func (g *Group) Wait() error
// Wait 会阻塞直至由上述 Go 方法调用的所有函数都返回，然后从它们返回第一个非nil的错误（如果有）

func fetchUrlDemo() error {
	g := new(errgroup.Group) //创建等待组，类似sync.WaitGroup
	var urls = []string{
		"http://pkg.go.dev",
		"http://www.liwenzhou.com",
		"http://www.yixieqitawangzhi.com",
	}

	for _, url := range urls {
		url := url
		// 启动goroutine去获取Url内容
		g.Go(func() error {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("获取%s成功\n", url)
				resp.Body.Close()
			}
			return err //返回错误
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("所有的groutine均成功")
	return nil
}
func main() {
	fetchUrlDemo()
}
