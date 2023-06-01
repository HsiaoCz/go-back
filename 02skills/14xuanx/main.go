package main

type OptFunc func(*Opts)

type Opts struct {
	maxConn int
	id      string
	tls     bool
}

type Server struct {
	opts Opts
}

func defaultOpts() Opts {
	return Opts{
		maxConn: 10,
		id:      "default",
		tls:     false,
	}
}

func withTLS(isTls bool) OptFunc {
	return func(o *Opts) {
		o.tls = isTls
	}
}

func withID(id string) OptFunc {
	return func(o *Opts) {
		o.id = id
	}
}

func withMaxConn(maxconn int) OptFunc {
	return func(o *Opts) {
		o.maxConn = maxconn
	}
}

func newServer(opts ...OptFunc) *Server {
	o := defaultOpts()

	for _, fn := range opts {
		fn(&o)
	}

	return &Server{
		opts: o,
	}
}
