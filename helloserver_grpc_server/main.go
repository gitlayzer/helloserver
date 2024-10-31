package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gitlayzer/helloserver/helloserver"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	listenAddr string
	listenPort string
)

type server struct {
	helloserver.UnimplementedGreeterServer
}

func (s *server) SayHello(_ context.Context, in *helloserver.HelloRequest) (*helloserver.HelloReply, error) {
	return &helloserver.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func init() {
	flag.StringVar(&listenAddr, "listen-addr", "0.0.0.0", "listen address")
	flag.StringVar(&listenPort, "listen-port", "8001", "listen port")
}

func main() {
	flag.Parse()

	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%s", listenAddr, listenPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	helloserver.RegisterGreeterServer(srv, &server{})

	if err = srv.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
