package controller

import (
	"net/http"
	"strconv"

	"github.com/berkkaradalan/GoRedisCache/bootstrap"
	"github.com/berkkaradalan/GoRedisCache/domain"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type TeamsController struct{
	TeamsUseCase 			domain.TeamsUseCase
	Env 					*bootstrap.Env
}

func (tc *TeamsController) GetTeamById(c *gin.Context, collection *mongo.Collection, redisClient *redis.Client) {
	team_id := c.Param("id")

	team, err := tc.TeamsUseCase.GetTeamById(c, collection, redisClient,team_id)

	if err != nil {
		c.JSON(http.StatusOK, domain.ErrAreaFetch)
		return
	}

	c.JSON(http.StatusOK, team)
	return
}

func (tc *TeamsController) GetTeams(c *gin.Context, collection *mongo.Collection) {
	offsetStr := c.DefaultQuery("offset", "0")
	limitStr := c.DefaultQuery("limit", "10")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset"})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}

	team, err := tc.TeamsUseCase.GetTeams(c, collection, offset, limit)

	if err != nil {
		c.JSON(http.StatusOK, domain.ErrAreaFetch)
		return
	}

	c.JSON(http.StatusOK, team)
	return
}