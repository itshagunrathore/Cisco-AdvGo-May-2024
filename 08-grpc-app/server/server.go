package main

import (
	"context"
	"grpc-app/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

// Service Implementation
type AppServiceImpl struct {
	// inherit proto.UnimplementedAppServiceServer
	proto.UnimplementedAppServiceServer
}

// Implementation of proto.AppServiceServer interface
func (asi *AppServiceImpl) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	x := req.GetX()
	y := req.GetY()
	log.Printf("[Add] processing x=%d and y=%d\n", x, y)
	result := x + y
	log.Printf("[Add] sending response\n")
	res := &proto.AddResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	asi := &AppServiceImpl{}
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grpcServer, asi)
	grpcServer.Serve(listener)
}
