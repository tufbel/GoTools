// Package grpc_clients
// Title       : grpc_clients.go
// Author      : Tuffy  2023/3/24 15:14
// Description :
package grpc_clients

import (
	gp "GoTools/grpc_eg/proto"
	"GoTools/tools/myLog"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientRun() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		myLog.Logger.Error(fmt.Sprintf("did not connect: %v", err))
		return
	}
	defer conn.Close()
	c := gp.NewHomeClient(conn)

	// 调用SayHello接口
	r, err := c.SayHello(context.Background(), &gp.HelloRequest{Name: "Tuffy"})
	if err != nil {
		myLog.Logger.Error(fmt.Sprintf("could not greet: %v", err))
		return
	}
	myLog.Logger.Debug(fmt.Sprintf("Greeting: %s", r.GetMessage()))

	// 调用SayBye接口
	r2, err := c.SayBye(context.Background(), &gp.ByeRequest{Name: "Jerry"})
	if err != nil {
		myLog.Logger.Error(fmt.Sprintf("could not say bye: %v", err))
		return
	}
	myLog.Logger.Debug(fmt.Sprintf("Saying bye: %s", r2.GetMessage()))
}
