package main

import (
	hello "animal/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"net/http"
)

var (
	client hello.HelloServiceClient
)

type MD map[string][]string

func myHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()

	fmt.Println("web ")
	label := r.Header.Get("label")
	code := r.Header.Get("code")
	md := metadata.New(map[string]string{"label": label, "code": code})
	ctx = metadata.NewOutgoingContext(ctx, md)
	pong, err := client.Ping(ctx, &hello.Req{})
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("err %s", err.Error()))
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("page %s", pong.Version))
}

func main() {
	http.HandleFunc("/", myHandler) //	设置访问路由
	fmt.Println("server start")
	conn, err := grpc.Dial("grpc-svc.default:50001", grpc.WithInsecure())
	//conn, err := grpc.Dial("localhost:50001", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client = hello.NewHelloServiceClient(conn)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
