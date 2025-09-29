package main

import (
	address_book "address-book/address-book"
	"address-book/config"
	"address-book/repository"
	"address-book/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	log.Println("Starting server on port 8080")

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Cannot create listener: %s", err)
	}

	log.Println("Started server: port 8080")

	mongoDbClient := config.SetUpMongo()

	// Set collections in the repository
	repository.SetCollections(mongoDbClient)

	server := grpc.NewServer()
	service := &service.NewAddressBook{}

	address_book.RegisterAddAddressBookServer(server, service)
	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("Cannot serve: %s", err)
	}
}
