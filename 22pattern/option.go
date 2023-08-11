package main

// 选项模式

type OptFunc func(opts *Opts)

type Opts struct {
	maxConn int
	id      string
	tls     bool
}

type server struct {
	Opts
}

func defaultOpts() Opts {
	return Opts{
		maxConn: 10,
		id:      "hello",
		tls:     false,
	}
}

func Withid(id string) OptFunc {
	return func(opts *Opts) {
		opts.id = id
	}
}

func WithMaxConn(maxConn int) OptFunc {
	return func(opts *Opts) {
		opts.maxConn = maxConn
	}
}

func Withtls(tls bool) OptFunc {
	return func(opts *Opts) {
		opts.tls = tls
	}
}

func NewServer(optf ...OptFunc) *server {
	opts := defaultOpts()

	for _, opt := range optf {
		opt(&opts)
	}

	return &server{Opts: opts}
}
