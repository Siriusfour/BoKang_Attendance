package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	MyProto "grpc-server/proto" // 这里根据你的实际路径修改
	"grpc-server/server"        // 你自己的 server 包路径
)

func main() {
	// 启动 TCP 监听
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	// 创建 gRPC 服务器
	grpcServer := grpc.NewServer()

	// 注册 LoginService 服务
	MyProto.RegisterLoginServiceServer(grpcServer, &server.Server{})

	// 启动服务
	fmt.Println("Server is listening on port :50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
