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
```


