package repository

import (
	"address-book/config"
	"log"
)

var mongoDbClient *config.AuditMongoClient

func SetCollections(mongoDbClient *config.AuditMongoClient) {
	// Set the MongoDB client for the repository
	SetMongoClient(mongoDbClient)

	addressCollection := mongoDbClient.Db.Collection("address")
	log.Println("Address collection created, addressCollection: ", addressCollection.Name())
}

func SetMongoClient(client *config.AuditMongoClient) {
	mongoDbClient = client
}
