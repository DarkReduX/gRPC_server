package main

import (
	"context"
	pb "github.com/DarkReduX/gRPC_service/protocol"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		logrus.Fatalf("Failed to listen on port 8080: %v", err)
	}

	_ = pb.UserNameMessage{}

	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, server{})
	if err = grpcServer.Serve(lis); err != nil {
		logrus.Fatalf("Failed to server gRPC server over port 8080")
	}

}

func (s server) SayHelloMessage(ctx context.Context, req *pb.UserNameMessage) (response *pb.HelloMessage, err error) {
	response = &pb.HelloMessage{
		Message: "Hello " + req.Name,
	}
	return response, nil
}
