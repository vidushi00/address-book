package main

import (
	address_book "address-book/address-book"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type newAddressBook struct {
	address_book.UnimplementedAddAddressBookServer
}

func (s *newAddressBook) AddAddress(ctx context.Context, req *address_book.AddAddressRequest) (*address_book.AddAddressResponse, error) {
	log.Printf("Added address successfully: %v", req)
	return &address_book.AddAddressResponse{Message: "Address added"}, nil
}

func main() {
	log.Println("Starting server on port 8080")

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Cannot create listener: %s", err)
	}

	log.Println("Started server: port 8080")

	server := grpc.NewServer()
	service := &newAddressBook{}

	address_book.RegisterAddAddressBookServer(server, service)
	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("Cannot serve: %s", err)
	}
}
