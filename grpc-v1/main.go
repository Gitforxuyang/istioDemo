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

var (
	client hello.HelloServiceClient
)

type HelloService struct {
}

func (h HelloService) Ping(ctx context.Context, req *hello.Req) (*hello.Pong, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		fmt.Println(md.Get("label"))
	}
	fmt.Println(md.Get("abc") == nil)
	ctx = metadata.NewOutgoingContext(ctx, md)
	//client.Ping(ctx, req)
	//labels := md.Get("label")
	//if len(labels) != 0 && labels[0] == "error" {
	//	return nil, errors.New("err")
	//}
	return &hello.Pong{Version: "v1"}, nil
}

func main() {
	grpcServer := grpc.NewServer()
	hello.RegisterHelloServiceServer(grpcServer, new(HelloService))

	lis, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := grpc.Dial("localhost:50002", grpc.WithInsecure())
	//conn, err = grpc.Dial("localhost:50002", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client = hello.NewHelloServiceClient(conn)
	fmt.Println("grpc v1 start")
	grpcServer.Serve(lis)
}
