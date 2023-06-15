package main

import "fmt"

// 选项模式

type OptionFunc func(*Opts)

type Opts struct {
	maxConn int
	isTls   bool
	id      string
}

type Server struct {
	Opts
}

func defaultOpts() Opts {
	return Opts{
		maxConn: 10,
		id:      "Hello",
		isTls:   false,
	}
}

func withMaxConn(maxConn int) OptionFunc {
	return func(o *Opts) {
		o.maxConn = maxConn
	}
}

func withId(id string) OptionFunc {
	return func(o *Opts) {
		o.id = id
	}
}

func withTls(isTls bool) OptionFunc {
	return func(o *Opts) {
		o.isTls = isTls
	}
}

func newServer(opts ...OptionFunc) *Server {
	o := defaultOpts()

	for _, fn := range opts {
		fn(&o)
	}
	return &Server{
		Opts: o,
	}
}

func main() {
	s := newServer(withId("hi"), withMaxConn(100), withTls(true))
	fmt.Printf("%+v\n", s)
}
