package main

import (
	"context"
	"fmt"
	pb "github.com/rzmsv/grpc-golang/proto/person"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient("localhost:8080", opts...)
	if err != nil {
		log.Fatalf("Can not connect client server: %v", err)
	}
	defer conn.Close()
	client := pb.NewPersonServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	/* ---------------------------- CREATE NEW PERSON --------------------------- */
	person := &pb.CreatePersonRequest{
		Name:        "Reza mousavi",
		Email:       "Rmussavi@gmail.com",
		PhoneNumber: "09148876040",
	}
	resPerson, err := client.Create(ctx, person)
	if err != nil {
		log.Fatalf("Can not create new user: %v", err)
	}
	fmt.Println(resPerson)
}
