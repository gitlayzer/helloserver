# HelloServer

## 介绍
```text
Hello Server 是一个基于 Golang 实现的基础的模块调用服务, 它使用了 gRPC 作为模块间的通信协议, 并提供了 HTTP 接口对外提供服务.
``` 

## 功能
- 模块内使用的是 gRPC 进行通信, gRPC Server 默认监听 8001 端口, 可以使用 --help 查看 flag 参数
- HTTP Server 对外提供了 Restful API 接口, 用户可以使用 HTTP 方式请求 HTTP Server, 内部走 gRPC 协议

## 示例
##### 注意：此项目 Go 版本为 1.22.8
```text
1. 克隆项目
   $ git clone https://github.com/gitlayzer/helloserver.git
   
2. 启动 gRPC Server
   $ go run helloserver_grpc_server\main.go

3. 启动 HTTP Server
   $ go run helloserver_http_server\main.go

4. 测试请求 HTTP Server
   $ curl http://localhost:8000/hello?name=gitlayzer
     {
       "code": 200,
       "msg": "Hello gitlayzer",
       "timestamp": "2024-10-30 19:36:02",
       "uuid": "e38cae8c-b780-4775-9279-9375687b8435"
     }
```