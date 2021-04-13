package main

import (
	hello "animal/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"math/rand"
	"net/http"
)

var (
	client  hello.HelloServiceClient
	client2 hello.HelloServiceClient
)

type MD map[string][]string

func myHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()

	fmt.Println("web ")
	label := r.Header.Get("label")
	code := r.Header.Get("code")
	fmt.Println(r.Header)
	md := metadata.New(map[string]string{"label": label, "code": code})
	ctx = metadata.NewOutgoingContext(ctx, md)
	//md = metadata.New(map[string]string{"flag": "true"})
	//ctx = metadata.NewOutgoingContext(ctx, md)
	if rand.Int31n(10) < 5 {
		pong, err := client.Ping(ctx, &hello.Req{})
		if err != nil {
			fmt.Fprintf(w, fmt.Sprintf("err %s", err.Error()))
			return
		}
		fmt.Fprintf(w, fmt.Sprintf("page 1 %s", pong.Version))
	} else {

		pong, err := client2.Ping(ctx, &hello.Req{})
		if err != nil {
			fmt.Fprintf(w, fmt.Sprintf("err %s", err.Error()))
			return
		}
		fmt.Fprintf(w, fmt.Sprintf("page 2 %s", pong.Version))
	}
}

func main() {
	http.HandleFunc("/", myHandler) //	设置访问路由
	fmt.Println("server start")
	conn, err := grpc.Dial("grpc-svc.default:50001", grpc.WithInsecure())
	//conn, err := grpc.Dial("localhost:50001", grpc.WithInsecure())
	//conn, err = grpc.Dial("localhost:50001", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client = hello.NewHelloServiceClient(conn)
	conn2, err := grpc.Dial("grpc-svc.default:50001", grpc.WithInsecure())
	client2 = hello.NewHelloServiceClient(conn2)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
