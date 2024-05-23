package main

import (
	"context"
	"grpc-app/proto"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

// Service Implementation
type AppServiceImpl struct {

	// inherit proto.UnimplementedAppServiceServer
	proto.UnimplementedAppServiceServer
}

// Implementation of proto.AppServiceServer interface
// Overriding the Add() method of "UnimplementedAppServiceServer"
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

func (asi *AppServiceImpl) GeneratePrimes(req *proto.PrimeRequest, serverStream proto.AppService_GeneratePrimesServer) error {
	start := req.GetStart()
	end := req.GetEnd()
	log.Printf("[GeneratePrimes] Received req for generating primes from %d to %d\n", start, end)
	for no := start; no <= end; no++ {
		if isPrime(no) {
			res := &proto.PrimeResponse{
				PrimeNo: no,
			}
			log.Printf("[GeneratePrimes] Sending prime no : %d\n", no)
			if err := serverStream.Send(res); err != nil {
				log.Fatalln(err)
			}
			time.Sleep(500 * time.Millisecond)
		}
	}
	return nil
}

func isPrime(no int64) bool {
	for i := int64(2); i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
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
