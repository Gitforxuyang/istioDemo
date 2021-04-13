package main

import (
	"animal/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
)

type HelloService struct {
}

func (h HelloService) Ping(ctx context.Context, req *hello.Req) (*hello.Pong, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		fmt.Println(md.Get("label"))
	}
	fmt.Println(md)
	return &hello.Pong{Version: "v2"}, nil
}

func main() {
	grpcServer := grpc.NewServer()
	hello.RegisterHelloServiceServer(grpcServer, new(HelloService))

	lis, err := net.Listen("tcp", ":50002")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("grpc v2 start")
	grpcServer.Serve(lis)
}
