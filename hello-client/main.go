package main

import (
	"context"
	"fmt"
	pb "gRPC_Study/hello-server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//通过grpc连接到server端，此处禁用安全传输，没有加密验证
	conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("failed to connect: %v", err)
	}

	defer conn.Close()

	//创建客户端建立连接
	client := pb.NewSayHelloClient(conn)

	//执行rpc调用，在服务端实现具体逻辑并返回结果
	client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "world!"})
}
