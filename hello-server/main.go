package main

import (
	"context"
	"fmt"
	pb "gRPC_Study/hello-server/proto"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	pb.UnimplementedSayHelloServer
}

func (s server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{ResponseMsg: "hello" + req.RequestName}, nil
}

func main() {
	//开启端口
	listen, _ := net.Listen("tcp", ":8888")
	//创建grpc服务
	grpcServer := grpc.NewServer()
	//注册自己编写的服务
	pb.RegisterSayHelloServer(grpcServer, &server{})

	//启动服务
	err := grpcServer.Serve(listen)
	if err != nil {
		fmt.Println("failed to serve: ", err)
		return
	}
}
