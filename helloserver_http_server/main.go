package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gitlayzer/helloserver/helloserver"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

var (
	listenAddr  string
	listenPort  string
	environment string
	grpcAddr    string
	grpcPort    string
)

func SayHello(c *gin.Context) {
	var name = c.Query("name")

	requestID, _ := uuid.NewRandom()

	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", grpcAddr, grpcPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(500, gin.H{
			"code":      500,
			"msg":       "grpc conn err",
			"uuid":      requestID,
			"timestamp": time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	client := helloserver.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.SayHello(ctx, &helloserver.HelloRequest{Name: name})
	if err != nil {
		c.JSON(500, gin.H{
			"code":      500,
			"msg":       "grpc call err",
			"uuid":      requestID,
			"timestamp": time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	c.JSON(200, gin.H{
		"code":      200,
		"msg":       r.GetMessage(),
		"uuid":      requestID,
		"timestamp": time.Now().Format("2006-01-02 15:04:05"),
	})
}

func init() {
	flag.StringVar(&grpcAddr, "grpc-addr", "localhost", "grpc address")
	flag.StringVar(&grpcPort, "grpc-port", "8001", "grpc port")
	flag.StringVar(&listenAddr, "listen-addr", "0.0.0.0", "listen address")
	flag.StringVar(&listenPort, "listen-port", "8000", "listen port")
	flag.StringVar(&environment, "environment", "debug", "environment")
}

func main() {
	flag.Parse()

	gin.SetMode(environment)

	r := gin.Default()

	r.GET("/hello", SayHello)

	if err := r.Run(fmt.Sprintf("%s:%s", listenAddr, listenPort)); err != nil {
		log.Fatalln(err)
	}
}
