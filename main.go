package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/micnncim/grpc-go-sample/proto"
)

type GreetingService struct{}

func (s *GreetingService) Hello(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{Message: "Hello, " + req.Name + "!"}, nil
}

func (s *GreetingService) Bye(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{Message: "Bye, " + req.Name + "!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	pb.RegisterGreetingServiceServer(server, &GreetingService{})
	reflection.Register(server)
	server.Serve(lis)
}
