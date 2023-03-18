// Package grpc_servers
// Title       : grpc_services.go
// Author      : Tuffy  2023/3/11 14:16
// Description :
package grpc_servers

import (
	gp "GoTools/grpc_eg/proto"
	"GoTools/tools/myLog"
	"context"
	"fmt"
)

type HomeServer struct {
	gp.UnimplementedHomeServer
}

func (s *HomeServer) SayHello(ctx context.Context, in *gp.HelloRequest) (*gp.HelloReply, error) {
	myLog.Logger.Debug(fmt.Sprintf("SayHello received: %v", in.GetName()))
	return &gp.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *HomeServer) SayBye(ctx context.Context, in *gp.ByeRequest) (*gp.ByeReply, error) {
	myLog.Logger.Debug(fmt.Sprintf("SayBye received: %v", in.GetName()))
	return &gp.ByeReply{Message: "Bye " + in.GetName()}, nil
}
