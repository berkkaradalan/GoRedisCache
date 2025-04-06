package repository

import (
	"github.com/berkkaradalan/GoRedisCache/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlayersRepository interface {
	GetPlayerById(ctx *gin.Context, collection *mongo.Collection, player_id string) (*domain.PlayerResponse, error)
}

type playersRepository struct {
	mongodb *mongo.Collection
}

func NewPlayerRepository(mongo *mongo.Collection) PlayersRepository {
	return &playersRepository{
		mongodb: mongo,
	}
}

func (r *playersRepository) GetPlayerById(ctx *gin.Context, collection *mongo.Collection, player_id string) (*domain.PlayerResponse, error) {
	filter := bson.M{"_id": player_id}
	var player *domain.PlayerResponse
	err := collection.FindOne(ctx, filter).Decode(player)
	if err != nil {
		return nil, err
	}

	return player, nil
}