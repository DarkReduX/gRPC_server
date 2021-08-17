package main

import (
	"context"
	pb "github.com/DarkReduX/gRPC_service/protocol"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"io"
	"net"
	"os"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

type pictureServer struct {
	grpc.ServerStream
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

func (s server) SayHello(ctx context.Context, req *pb.UserNameMessage) (response *pb.HelloMessage, err error) {
	response = &pb.HelloMessage{
		Message: "Hello " + req.Name,
	}
	return response, nil
}

func (s server) SendPicture(stream pb.HelloService_SendPictureServer) error {
	pictureOut, err := os.Create("pic1.png")
	if err != nil {
		return err
	}
	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			responseMsg := "New responseMsg received "
			stream.SendAndClose(&pb.Response{Message: responseMsg})
			return nil
		}
		if err != nil {
			logrus.Fatalf("couldn't receive from stream : %v", err)
		}
		pictureOut.Write(msg.GetContent())
	}
}
