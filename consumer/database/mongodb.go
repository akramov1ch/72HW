package database

import (
	"72HW/consumer/model"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var ordersCollection *mongo.Collection

func init() {
	var err error
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	ordersCollection = client.Database("orderdb").Collection("orders")
}

func SaveOrder(order model.Order) error {
	filter := bson.M{"id": order.ID}
	update := bson.M{"$set": order}
	_, err := ordersCollection.UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))
	return err
}
