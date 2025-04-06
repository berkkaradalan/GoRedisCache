package controller

import (
	"net/http"

	"github.com/berkkaradalan/GoRedisCache/bootstrap"
	"github.com/berkkaradalan/GoRedisCache/domain"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlayerController struct{
	PlayerUseCase 			domain.PlayerUseCase
	Env 					*bootstrap.Env
}

func (pc *PlayerController) GetPlayerById(c *gin.Context, collection *mongo.Collection, redisClient *redis.Client) {
	player_id := c.Param("id")

	player, err := pc.PlayerUseCase.GetPlayerById(c, collection, redisClient,player_id)

	if err != nil {
		c.JSON(http.StatusOK, domain.ErrPlayerFetch)
		return
	}

	c.JSON(http.StatusOK, player)
	return
}