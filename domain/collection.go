package domain

import "go.mongodb.org/mongo-driver/mongo"

type MongoDBCollections struct{
	Areas   *mongo.Collection
	Teams   *mongo.Collection
	Players *mongo.Collection
	Matches *mongo.Collection
}