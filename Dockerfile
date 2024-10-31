# 定义基础构建镜像
FROM golang:1.22.8-alpine as builder
# 定义工作目录
WORKDIR /app
# 设置 Go 代理
ENV GOPROXY=https://goproxy.cn,direct
# 复制依赖文件
COPY go.mod go.sum ./
# 复制服务代码
COPY . .
# 构建可执行文件,使用 ARG 定义的服务名称
ARG NAME
RUN CGO_ENABLED=0 go build -ldflags -o ${NAME} ${NAME}/main.go

# 定义运行镜像
FROM busybox as runner
# 复制构建好的二进制文件
ARG NAME
COPY --from=builder /app/${NAME} /${NAME}
# 设置入口点
ENTRYPOINT ["/${NAME}"]