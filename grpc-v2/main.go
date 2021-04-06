package main

import (
	"animal/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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
	//time.Sleep(time.Second * 2)
	if len(labels) != 0 && labels[0] == "error" {
		return nil, status.Error(codes.Internal, "内部错误")
	}
	return &hello.Pong{Version: "v2"}, nil
}

func main() {
	grpcServer := grpc.NewServer()
	hello.RegisterHelloServiceServer(grpcServer, new(HelloService))

	lis, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("grpc v2 start")
	grpcServer.Serve(lis)
}
