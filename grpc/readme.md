## 1、protobuf

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-gr
pc_opt=paths=source_relative pb/greeter.proto

protoc -I=v1 \ master ✱
--go_out=v1 --go_opt=paths=source_relative \
 --go-grpc_out=v1 --go-grpc_opt=paths=source_relative \
 --grpc-gateway_out=v1 --grpc-gateway_opt=paths=source_relativ
e \
 v1/greeter.proto

### metadata

```go
// matedata 生成
// 第一种方式
md :=matedata.New(map[string]string{"key1":"v1","key2":"v2"})
// 第二种方式，key
md:=metadata.Pairs(
    "key1","val1",
    "key1","val2",
    "key2","val2",
)

// 发送matedata
// 新建一个matedata的context
ctx:=metadata.NewOutgoingContext(context.Background(),md)

// 接收metadata
md,ok:=metadata.FromIncomingContext(ctx)
```

### 拦截器

类似于中间件，我们实现拦截器，然后配置到项目里面就好了

可以看看 go-grpc-middleware

里面有一些常用的 grpc 中间件，都是通过实现拦截器实现的

### 拦截器和 metadata 实现 token 机制

### grpc 的验证器

这里要去 github 上整这个 proto-gen-validate

```go
protoc -I=. --go_out=plugins=grpc:. --validate_out="lang=go:." pb/ hello.proto
```

我测，骚报错
