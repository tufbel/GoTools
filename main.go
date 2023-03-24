package main

import (
	"GoTools/grpc_eg/grpc_clients"
	"GoTools/grpc_eg/grpc_servers"
	gp "GoTools/grpc_eg/proto"
	_ "GoTools/struct_and_class"
	"GoTools/tools/myLog"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"time"
)

func startGrpcServer() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		myLog.Logger.Error(fmt.Sprintf("SayHello received: %v", err))
		return
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	gp.RegisterHomeServer(grpcServer, &grpc_servers.HomeServer{})
	grpcServer.Serve(lis)
}

func main() {
	go startGrpcServer()
	time.Sleep(5 * time.Second)
	go grpc_clients.ClientRun()
	time.Sleep(100 * time.Second)
}
