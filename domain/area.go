package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type Area struct {
	ID					int				`bson:"id"`
	Name				string			`bson:"name"`
	CountryCode			string			`bson:"countryCode"`
	Flag				*string			`bson:"flag"`
	ParentAreaID		int				`bson:"parentAreaId"`
	ParentArea			string			`bson:"parentArea"`
	ChildAreas			[]*Area			`bson:"childAreas"`
}

type AreaResponse struct {
	ID					int				`bson:"id"`
	Name				string			`bson:"name"`
	CountryCode			string			`bson:"countryCode"`
	Flag				*string			`bson:"flag"`
	ParentAreaID		int				`bson:"parentAreaId"`
	ParentArea			string			`bson:"parentArea"`
	ChildAreas			[]*Area			`bson:"childAreas"`
}

type AreaUseCase interface {
	GetAreaById(ctx *gin.Context, collection *mongo.Collection, redisClient *redis.Client, area_id string) (AreaResponse, error)
}