package repository

import (
	"address-book/config"
	"address-book/model"
	"context"
)

var AddressCollection = "address"

func AddAddress(address model.Address, ctx context.Context) error {

	addressCollection := config.MongoDb.Collection(AddressCollection)
	_, err := addressCollection.InsertOne(ctx, address)
	if err != nil {
		return err
	}
	return nil

}
