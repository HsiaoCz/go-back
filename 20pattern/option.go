package main

// 选项模式

type OptFunc func(opts *Opts)

type Opts struct{
    tls bool
    id string
    maxConn int
}

type server struct{
    Opts
}

func defaultOpts()Opts{
    return Opts{
        maxConn:10,
        id:"hello",
        tls:false,
    }
}

func WithMaxConn(maxConn int)OptionFunc{
    return func(o *Opts){
        o.maxConn=maxConn
    }
}

func Withtls(tls bool)OptionFunc{
    return func(o *Opts){
        o.maxConn=maxConn
    }
}

func Withid(id string)OptionFunc{
    return func(o *Opts){
        o.id=id
    }
}

func NewServer(optfuncs ...OptionFunc) *server {
	o := defaultOpts()

	for _, optfunc := range optfuncs {
		optfunc(&o)
	}
	return &server{Opts: o}
}

