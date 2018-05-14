package main

import (
	"context"
	"fmt"
	"log"

	pb "room-one/proto"

	"google.golang.org/grpc"
)

const (
	address = "localhost:3000"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewEncryptionServiceClient(conn)

	fmt.Println("Enter the message to be reversed then encrypted by AES128")
	var text string
	fmt.Scanln(&text)
	reversed, err := client.Reverse(context.Background(), &pb.InputString{InitalInput: text})
	if err != nil {
		log.Fatalf("Could not encrypt", err)
	}
	fmt.Println("Reversed message: " + reversed.FinalOutput)

	response, err := client.Encrypt(context.Background(), &pb.RequestText{PlainMessage: reversed.FinalOutput})
	fmt.Println("Encrypted message: " + response.EncryptedMessage)
}
