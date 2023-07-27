package main

import "fmt"

type OptFunc func(*Opts)

type Opts struct {
	maxConn int
	id      string
	tls     bool
}

type Server struct {
	Opts
}

func defaultOpts() Opts {
	return Opts{
		maxConn: 10,
		id:      "default",
		tls:     false,
	}
}

func withTLS(isTLS bool) OptFunc {
	return func(o *Opts) {
		o.tls = isTLS
	}
}

func withMaxConn(maxConn int) OptFunc {
	return func(o *Opts) {
		o.maxConn = maxConn
	}
}

func withID(id string) OptFunc {
	return func(o *Opts) {
		o.id = id
	}
}

func newServer(opts ...OptFunc) *Server {
	o := defaultOpts()

	for _, fn := range opts {
		fn(&o)
	}
	return &Server{
		Opts: o,
	}
}
func main() {
	s := newServer(withTLS(false), withMaxConn(99), withID("hello"))
	fmt.Printf("%+v\n", s)
}
