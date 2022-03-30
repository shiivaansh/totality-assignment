package Mongo

import "go.mongodb.org/mongo-driver/mongo"

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("Totality Corp").Collection(collectionName)
	return collection
}
