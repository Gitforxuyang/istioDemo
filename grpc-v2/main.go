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
	"strconv"
	"time"
)

type HelloService struct {
}

func (h HelloService) Ping(ctx context.Context, req *hello.Req) (*hello.Pong, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		fmt.Println(md.Get("label"))
	}
	fmt.Println(md)
	//labels := md.Get("label")
	code := md.Get("code")
	//time.Sleep(time.Second * 2)
	if len(code) != 0 {
		fmt.Println(code, time.Now())
		c, _ := strconv.ParseInt(code[0], 10, 64)
		return nil, status.Error(codes.Code(c), "内部错误")
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
