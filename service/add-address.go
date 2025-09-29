package service

import (
	address_book "address-book/address-book"
	"address-book/model"
	"address-book/repository"
	"context"
	"log"
)

type NewAddressBook struct {
	address_book.UnimplementedAddAddressBookServer
}

func (s *NewAddressBook) AddAddress(ctx context.Context, req *address_book.AddAddressRequest) (*address_book.AddAddressResponse, error) {
	log.Printf("Added address successfully: %v", req)

	address := model.Address{
		Name:    req.Name,
		Email:   req.Email,
		Phone:   req.Phone,
		Address: req.Address,
	}
	err := repository.AddAddress(address, ctx)
	if err != nil {
		return nil, err
	}
	return &address_book.AddAddressResponse{Message: "Address added"}, nil
}
