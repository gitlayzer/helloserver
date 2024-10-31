FROM golang:1.22.8-alpine as builder
ARG TARGET
ENV GOPROXY=https://goproxy.cn,direct
WORKDIR /app
COPY go.mod go.sum ./
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go mod download && go build -ldflags="-s -w" -o bin/${TARGET} ./${TARGET}/main.go

FROM alpine:latest as runner
ARG TARGET
COPY --from=builder /app/bin/${TARGET} /app
ENTRYPOINT ["/app"]
