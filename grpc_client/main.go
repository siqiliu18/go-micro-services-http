package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	grpcproto "grpcclient/proto"
)

// Client is like the Service struct in RBK
type Client struct {
	chatServiceClient grpcproto.ChatServiceClient
	sayHelloRes       string
}

// GuiFunc is to show messages on UI
func (c *Client) GuiFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Say Hi!</h1>")
	fmt.Fprintf(w, "<p>"+c.sayHelloRes+"</p>")
}

func main() {
	conn, err := grpc.Dial("chat-server-service:8889", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	clientService := Client{
		chatServiceClient: grpcproto.NewChatServiceClient(conn),
	}

	// c := grpcproto.NewChatServiceClient(conn)

	res, err := clientService.chatServiceClient.SayHello(context.Background(), &grpcproto.Message{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", res.GetBody())
	clientService.sayHelloRes = res.GetBody()

	http.HandleFunc("/chat", clientService.GuiFunc)
	fmt.Println("Application is running on port: 9001")
	http.ListenAndServe(":9001", nil)
}
