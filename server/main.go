package main

import (
	"context"
	"errors"
	"log"
	"net"

	pb "github.com/rzmsv/grpc-golang/proto/person"
	"google.golang.org/grpc"
)

type Person struct {
	ID          int32
	Name        string
	Email       string
	PhoneNumber string
}

type server struct {
	pb.UnimplementedPersonServiceServer
}

var nextId int32 = 1
var persons = make(map[int32]Person)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Can not listen port: %v", err)
	}
	defer lis.Close()
	grpcServer := grpc.NewServer()
	pb.RegisterPersonServiceServer(grpcServer, &server{})
	log.Printf("gRPC server connect on port %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Can not serv port: %v", err)
	}
}

func (s *server) Create(ctx context.Context, in *pb.CreatePersonRequest) (*pb.PersonProfileResponse, error) {
	person := Person{
		Name:        in.GetName(),
		Email:       in.GetEmail(),
		PhoneNumber: in.GetPhoneNumber(),
	}

	if person.Name == "" || person.Email == "" || person.PhoneNumber == "" {
		return &pb.PersonProfileResponse{}, errors.New("Fields missing!")
	}
	person.ID = nextId
	persons[nextId] = person
	nextId++

	return &pb.PersonProfileResponse{
		Id:          person.ID,
		Name:        person.Name,
		Email:       person.Email,
		PhoneNumber: person.PhoneNumber,
	}, nil

}
