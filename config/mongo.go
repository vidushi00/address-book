package config

import (
	"address-book/dto"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDb *mongo.Database

type AuditMongoClient struct {
	Client *mongo.Client
	Db     *mongo.Database
}

func SetUpMongo() *AuditMongoClient {
	ctx := context.TODO()
	log.Println(" ⛁ Connecting to MongoDB Database")

	configuration := dto.Configuration{
		Mongo: dto.MongoConfiguration{
			UserName:     "root",
			Password:     "password",
			Url:          "localhost:27017", // or "127.0.0.1:27017" if localhost doesn't work
			AuthDatabase: "admin",
			Database:     "address-book",
		},
	}
	uri := fmt.Sprintf("mongodb://%s:%s@%s/?authSource=%s", configuration.Mongo.UserName,
		configuration.Mongo.Password, configuration.Mongo.Url, configuration.Mongo.AuthDatabase)
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("⛒ Connection Failed to Database, error: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("⛒ Connection Failed to Database, error: %v", err)
	}

	log.Println("⛁ Connected to MongoDB Database")
	MongoDb = client.Database(configuration.Mongo.Database)

	AuditMongoDb := &AuditMongoClient{Client: client, Db: MongoDb}
	return AuditMongoDb
}
