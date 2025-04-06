package repository

import (
	"github.com/berkkaradalan/GoRedisCache/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AreasRepository interface {
	GetAreaById(ctx *gin.Context, collection *mongo.Collection, area_id string) (*domain.AreaResponse, error)
}

type areasRepository struct {
	mongodb *mongo.Collection
}

func NewAreasRepository(mongo *mongo.Collection) AreasRepository {
	return &areasRepository{
		mongodb: mongo,
	}
}

func (r *areasRepository) GetAreaById(ctx *gin.Context, collection *mongo.Collection, area_id string) (*domain.AreaResponse, error) {
    filter := bson.M{"_id": area_id}
    var area *domain.AreaResponse
    err := collection.FindOne(ctx, filter).Decode(area)
    if err != nil {
        return nil, err
    }

    return area, nil
}
