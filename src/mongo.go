package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var collection *mongo.Collection

func initMongoDBConnection() (*mongo.Client, *mongo.Collection) {
	fmt.Println("Connecting to the mongodb database...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		panic(err)
	}
	/*
		    defer func() {
				if err = client.Disconnect(ctx); err != nil {
					panic(err)
				}
			}()
	*/
	collection = client.Database("tt").Collection("tests")

	return client, collection
}

func insertOneDocument(collection *mongo.Collection, document interface{}) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, document)
	if err != nil {
		panic(err)
	}
}

func disconnectFromMongoDB(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
}
