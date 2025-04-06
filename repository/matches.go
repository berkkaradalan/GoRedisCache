package repository

import (
	"log"

	"github.com/berkkaradalan/GoRedisCache/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MatchesRepository interface{
	GetTeamMatches(ctx *gin.Context, collection *mongo.Collection, limit int, offset int, teamID string) ([]domain.MatchResponse, error)
}

type matchesRepository struct {
	mongodb *mongo.Collection
}

func NewMatchesRepository(mongo *mongo.Collection) MatchesRepository {
	return &matchesRepository{
		mongodb: mongo,
	}
}

func (r *matchesRepository) GetTeamMatches(ctx *gin.Context, collection *mongo.Collection, limit int, offset int, teamID string) ([]domain.MatchResponse, error){
	var matches []domain.MatchResponse
	cursor, err:= collection.Find(ctx, bson.M{}, options.Find().SetSkip(int64(offset)).SetLimit(int64(limit)))

	if err!=nil {
		log.Println("Mongodb Find Error : %v", err)
		return nil, err
	}

	if err := cursor.Err(); err!=nil {
		log.Println("Cursor error : %v", err)
		return nil, err
	}

	for cursor.Next(ctx){
		var match domain.MatchResponse
		err:=cursor.Decode(&match)
		if err != nil {
			return nil, err
		}
		matches = append(matches, match)
	}

	return matches, nil
}