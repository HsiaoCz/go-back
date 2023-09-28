package main

// 选项模式

type Opts struct {
	maxConn int
	id      string
	tls     bool
}

type OptFunc func() *Opts

type Server struct {
	Opts
}

func WithMaxConn(maxConn int) OptFunc {
	return func() *Opts {
		return &Opts{
			maxConn: maxConn,
		}
	}
}

func WithId(id string) OptFunc {
	return func() *Opts {
		return &Opts{
			id: id,
		}
	}
}

func WithTls(tls bool) OptFunc {
	return func() *Opts {
		return &Opts{
			tls: tls,
		}
	}
}

func DefaultOpts() Opts {
	return Opts{
		maxConn: 10,
		id:      "helloo",
		tls:     false,
	}
}

func NewServer(optFuncs ...OptFunc) *Server {
	o := DefaultOpts()
	for _, optfunc := range optFuncs {
		optfunc()
	}
	return &Server{
		Opts: o,
	}
}
