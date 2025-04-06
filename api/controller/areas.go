package controller

import (
	"github.com/berkkaradalan/GoRedisCache/bootstrap"
	"github.com/berkkaradalan/GoRedisCache/domain"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type AreaController struct{
	AreaUseCase 			domain.AreaUseCase
	Env 					*bootstrap.Env
}

func (ac *AreaController) GetAreaById(c *gin.Context, collection *mongo.Collection, redisClient *redis.Client) {
	area_id := c.Param("id")

	area, err := ac.AreaUseCase.GetAreaById(c, collection, redisClient,area_id)

	if err != nil {
		c.JSON(http.StatusOK, domain.ErrAreaFetch)
		return
	}

	c.JSON(http.StatusOK, area)
	return
}