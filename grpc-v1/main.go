package main

import (
	"animal/proto"
	"context"
	"errors"
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
	labels := md.Get("label")
	if len(labels) != 0 && labels[0] == "error" {
		return nil, errors.New("err")
	}
	return &hello.Pong{Version: "v1"}, nil
}

func main() {
	grpcServer := grpc.NewServer()
	hello.RegisterHelloServiceServer(grpcServer, new(HelloService))

	lis, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("grpc v1 start")
	grpcServer.Serve(lis)
}
