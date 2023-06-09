package main

import (
	"context"
	"log"
	"tutorial-gate/chat"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to 9000!, %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "Hello from the client!",
	}
	res, err := c.SayHello(context.Background(), &message)

	if err != nil {
		log.Fatalf("Error when calling SayHello! %s", err)
	}

	log.Printf("Response from server: %s", res.Body)
}
