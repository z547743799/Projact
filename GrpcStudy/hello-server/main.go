package main

import (
	pb "GrpcStudy/hello-server/proto"
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	fmt.Println("hello" + req.Name)
	return &pb.HelloResponse{Message: "hello" + req.Name}, nil
}
func main() {
	//开启端口
	listen, _ := net.Listen("tcp", ":9090")
	//创建grpc服务
	grpcServer := grpc.NewServer()
	//在grpc服务端中去注册我们自己编写的服务
	pb.RegisterHelloServiceServer(grpcServer, &server{})
	//启动服务
	err := grpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("faild to servei:%v", err)
		return
	}
}
