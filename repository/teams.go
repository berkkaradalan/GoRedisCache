package repository

import (
	"log"

	"github.com/berkkaradalan/GoRedisCache/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TeamsRepository interface {
	GetTeamById(ctx *gin.Context, collection *mongo.Collection, team_id string) (*domain.TeamResponse, error)
	GetTeams(ctx *gin.Context, collection *mongo.Collection, limit int, offset int) ([]domain.TeamResponse, error)
}

type teamsRepository struct {
	mongodb *mongo.Collection
}

func NewTeamsRepository(mongo *mongo.Collection) TeamsRepository {
	return &teamsRepository{
		mongodb: mongo,
	}
}

func (r *teamsRepository) GetTeamById(ctx *gin.Context, collection *mongo.Collection, team_id string) (*domain.TeamResponse, error) {
    filter := bson.M{"_id": team_id}
	var team domain.TeamResponse
	err := collection.FindOne(ctx, filter).Decode(&team)
	if err != nil {
	    return nil, err
	}
	return &team, nil
}


func (r *teamsRepository) GetTeams(ctx *gin.Context, collection *mongo.Collection, limit int, offset int) ([]domain.TeamResponse, error) {
	var teams []domain.TeamResponse
	cursor, err := collection.Find(ctx, bson.M{}, options.Find().SetSkip(int64(offset)).SetLimit(int64(limit)))
	if err != nil {
		log.Printf("MongoDB Find error: %v", err)
		return nil, err
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Cursor error: %v", err)
		return nil, err
	}
    if err != nil {
        return nil, err
    }
	for cursor.Next(ctx) {
		var team domain.TeamResponse
		err := cursor.Decode(&team)
		if err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}
	return teams, nil
}