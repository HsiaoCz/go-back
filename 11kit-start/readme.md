## go-kit

GO kit 构建服务分为三层：

1、传输层(Trasnport)

传输域绑定到具体的传输协议，如 HTTP 或 gRPC。在一个微服务可能支持一个或多个传输协议的世界中，这是非常强大的：你可以在单个微服务中支持原有的 HTTP API 和新增的 RPC 服务

当实现 REST 式的 HTTP API 时，你的路由是在 HTTP 传输中定义的。最常见的路由定义在 HTTP 路由器函数中，如下所示:

```go
r.Methods("POST").Path("/profiles/").Handler(httptransport.NewServer(
		e.PostProfileEndpoint,
		decodePostProfileRequest,
		encodeResponse,
		options...,
))
```

2、端点层(Endpoint)

端点就像控制器上的动作/处理程序; 它是安全性和抗脆弱性逻辑的所在。如果实现两种传输(HTTP 和 gRPC) ，则可能有两种将请求发送到同一端点的方法

3、服务层(Service)

服务（指 Go kit 中的 service 层）是实现所有业务逻辑的地方。服务层通常将多个端点粘合在一起。在 Go kit 中，服务层通常被抽象为接口，这些接口的实现包含业务逻辑。Go kit 服务层应该努力遵守整洁架构或六边形架构。也就是说，业务逻辑不需要了解端点（尤其是传输域）概念：你的服务层不应该关心 HTTP 头或 gRPC 错误代码。

Middlewares

Go kit 试图通过使用中间件（或装饰器）模式来执行严格的关注分离（separation of concerns）。中间件可以包装端点或服务以添加功能，比如日志记录、速率限制、负载平衡或分布式跟踪。围绕一个端点或服务链接多个中间件是很常见的。

## 1、实例

新建一个 pb 文件夹，写一个 proto

生成文件

```bash
protoc -I=pb \
   --go_out=pb --go_opt=paths=source_relative \
   --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
   pb/kitstart.proto
```

新建三个文件夹，分别为 service transport endpoint

service 这层主要负责实现业务逻辑

```go
// service.go


```

endpoint 层负责存放项目对外暴露的 rpc 方法

```go
// endpoint.go
```

transport 表示项目对外通信的相关部分

```go

```

## 2、中间件

在 go kit 中，对中间件的定义是接收 Endpoint 并返回 Endpoint 的函数

```go
type Middleware func(Endpoint) Endpoint
```

```go
import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
)

func loggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			logger.Log("msg", "calling endpoint")
			defer logger.Log("msg", "called endpoint")
			return next(ctx, request)
		}
	}
}
```
