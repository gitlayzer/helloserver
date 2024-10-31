# HelloServer

## 介绍
```text
Hello Server 是一个基于 Golang 实现的基础的模块调用服务, 它使用了 gRPC 作为模块间的通信协议, 并提供了 HTTP 接口对外提供服务.
``` 

## 功能
- 模块内使用的是 gRPC 进行通信, gRPC Server 默认监听 8001 端口, 可以使用 --help 查看 flag 参数
- HTTP Server 对外提供了 Restful API 接口, 用户可以使用 HTTP 方式请求 HTTP Server, 内部走 gRPC 协议

## 示例

##### 项目 API 接口
> 注意：此项目 Go 版本为 1.22.8
```api
GET: /hello
Query Parameters:
  name: string
Response:
  {
    "code": int,
    "msg": string,
    "timestamp": string,
    "uuid": string
  }
```

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

## 构建 Docker 镜像

> 注：如果你的 Docker 是 18.x 或更高版本, 请安装 docker-buildx 插件, 并设置 DOCKER_BUILDKIT=1 环境变量

##### 构建 gRPC Server 镜像
```shell
$ docker build --rm --build-arg TARGET=helloserver_grpc_server -t helloserver_grpc_server:v0.0.1 .
```

##### 构建 HTTP Server 镜像
```shell
$ docker build --rm --build-arg TARGET=helloserver_http_server -t helloserver_http_server:v0.0.1 .
```

##### 运行 gRPC Server 容器
```shell
$ docker run -d --rm --name grpc-server -p 8001:8001 helloserver_grpc_server:v0.0.1
```

##### 运行 HTTP Server 容器
```shell
$ docker run -d --rm --name http-server -p 8000:8000 helloserver_http_server:v0.0.1 --grpc-addr <grpc-server-ip> --grpc-port <grpc-server-port>
```

## 其他
- 项目地址：https://github.com/gitlayzer/helloserver
- 项目作者：gitlayzer
- 作者邮箱：gduxintian@gmail.com