package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
LOOP:
	for no := start; no <= end; no++ {
		if isPrime(no) {
			res := &proto.PrimeResponse{
				PrimeNo: no,
			}
			log.Printf("[GeneratePrimes] Sending prime no : %d\n", no)
			if err := serverStream.Send(res); err != nil {
				// log.Fatalln(err)
				if code := status.Code(err); code == codes.Unavailable {
					fmt.Println("Transport closed")
					break LOOP
				}
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

func (as *AppServiceImpl) CalculateAverage(serverStream proto.AppService_CalculateAverageServer) error {
	var total, count int64
LOOP:
	for {
		req, err := serverStream.Recv()
		if err == io.EOF {
			avg := float32(total / count)
			fmt.Println("sending response, avg :", avg)
			res := &proto.AverageResponse{
				Average: avg,
			}
			serverStream.SendAndClose(res)
			break LOOP
		}
		if err != nil {
			log.Fatalln(err)
		}
		no := req.GetNo()
		fmt.Println("average req, no :", no)
		total += no
		count++
	}
	return nil
}

func (asi *AppServiceImpl) Greet(serverStream proto.AppService_GreetServer) error {
	for {
		greetReq, err := serverStream.Recv()
		if code := status.Code(err); code == codes.Unavailable {
			fmt.Println("Client connection closed")
			break
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		person := greetReq.GetPerson()
		firstName := person.GetFirstName()
		lastName := person.GetLastName()
		log.Printf("Received greet request for %q and %q\n", firstName, lastName)
		message := fmt.Sprintf("Hi %s %s, Have a nice day!", firstName, lastName)
		time.Sleep(2 * time.Second)
		log.Printf("Sending response : %q\n", message)
		greetResp := &proto.GreetResponse{
			Message: message,
		}
		if err := serverStream.Send(greetResp); err != nil {
			if code := status.Code(err); code == codes.Unavailable {
				fmt.Println("Client connection closed")
				break
			}
		}
	}
	return nil
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
