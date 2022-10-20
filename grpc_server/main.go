package main

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"

	"google.golang.org/grpc"

	grpcproto "grpcserver/proto"
)

// Server is a type of struct
type Server struct{}

// SayHello is a RPC function
func (s *Server) SayHello(ctx context.Context, in *grpcproto.Message) (*grpcproto.Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &grpcproto.Message{Body: "Hello From the Server!"}, nil
}

func main() {
	fmt.Println("Go gRPC Beginners Tutorial!")

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("falied to listen: %v", err)
	}

	svc := Server{}

	grpcServer := grpc.NewServer()

	grpcproto.RegisterChatServiceServer(grpcServer, &svc)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server: %s", err)
	}
}
