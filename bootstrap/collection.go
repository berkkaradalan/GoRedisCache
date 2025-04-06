package bootstrap

import (
	"github.com/berkkaradalan/GoRedisCache/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCollections(client *mongo.Client) *domain.MongoDBCollections {
	db := client.Database("football-data")
	return &domain.MongoDBCollections{
		Areas:   db.Collection("areas"),
		Teams:   db.Collection("teams"),
		Players: db.Collection("players"),
		Matches: db.Collection("matches"),
	}
}